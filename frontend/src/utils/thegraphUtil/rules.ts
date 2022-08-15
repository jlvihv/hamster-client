import { parseGRT } from '@graphprotocol/common-ts';
import gql from 'graphql-tag';
import { BigNumber } from 'ethers';
import {
  IndexerManagementClient,
  IndexingRuleAttributes,
  IndexingDecisionBasis,
} from '@graphprotocol/indexer-common';

export function pickFields(
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  rule: { [key: string]: any },
  keys: string[],
  drop: string[] = ['__typename'],
  // eslint-disable-next-line @typescript-eslint/ban-types
): object {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  let obj = {} as any;
  if (keys.length === 0) {
    obj = { ...rule };
  } else {
    for (const key of keys) {
      obj[key] = rule[key];
    }
  }
  for (const key of drop) {
    delete obj[key];
  }
  return obj;
}

export const parseBoolean = (val: string | boolean | number | undefined | null): boolean => {
  const s = val && val.toString().toLowerCase().trim();
  return s != 'false' && s != 'f' && s != '0';
};
export const parseDecisionBasis = (s: string): IndexingDecisionBasis => {
  if (!['always', 'never', 'rules', 'offchain'].includes(s)) {
    throw new Error(`Unknown decision basis "${s}". Supported: always, never, rules, offchain`);
  } else {
    return s as IndexingDecisionBasis;
  }
};

function nullPassThrough<T, U>(fn: (x: T) => U): (x: T | null) => U | null {
  return (x: T | null) => (x === null ? null : fn(x));
}
// eslint-disable-next-line @typescript-eslint/no-explicit-any
const INDEXING_RULE_PARSERS: Record<keyof IndexingRuleAttributes, (x: never) => any> = {
  id: (x) => x,
  identifier: (x) => x,
  identifierType: (x) => x,
  allocationAmount: nullPassThrough(parseGRT),
  allocationLifetime: nullPassThrough(parseInt),
  autoRenewal: nullPassThrough(parseBoolean),
  parallelAllocations: nullPassThrough(parseInt),
  minSignal: nullPassThrough(parseGRT),
  maxSignal: nullPassThrough(parseGRT),
  minStake: nullPassThrough(parseGRT),
  maxAllocationPercentage: nullPassThrough(parseFloat),
  minAverageQueryFees: nullPassThrough(parseGRT),
  decisionBasis: nullPassThrough(parseDecisionBasis),
  custom: nullPassThrough(JSON.parse),
  requireSupported: (x) => parseBoolean(x),
};
/**
 * Parses a user-provided indexing rule into a normalized form.
 */
export const parseIndexingRule = (
  rule: Partial<IndexingRuleAttributes>,
): Partial<IndexingRuleAttributes> => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const obj = {} as any;
  for (const [key, value] of Object.entries(rule)) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    obj[key] = (INDEXING_RULE_PARSERS as any)[key](value);
  }
  return obj as Partial<IndexingRuleAttributes>;
};
const INDEXING_RULE_CONVERTERS_TO_GRAPHQL: Record<
  keyof IndexingRuleAttributes,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  (x: never) => any
> = {
  id: (x) => x,
  identifier: (x) => x,
  identifierType: (x) => x,
  allocationAmount: nullPassThrough((x: BigNumber) => x.toString()),
  allocationLifetime: nullPassThrough((x: number) => x),
  autoRenewal: (x) => x,
  parallelAllocations: nullPassThrough((x: number) => x),
  minSignal: nullPassThrough((x: BigNumber) => x.toString()),
  maxSignal: nullPassThrough((x: BigNumber) => x.toString()),
  minStake: nullPassThrough((x: BigNumber) => x.toString()),
  maxAllocationPercentage: nullPassThrough((x: number) => x),
  minAverageQueryFees: nullPassThrough((x: BigNumber) => x.toString()),
  decisionBasis: (x) => x,
  custom: nullPassThrough(JSON.stringify),
  requireSupported: (x) => x,
};
/**
 * Converts a normalized indexing rule to a representation
 * compatible with the indexer management GraphQL API.
 */
export const indexingRuleToGraphQL = (
  rule: Partial<IndexingRuleAttributes>,
): Partial<IndexingRuleAttributes> => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const obj = {} as any;
  for (const [key, value] of Object.entries(rule)) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    obj[key] = (INDEXING_RULE_CONVERTERS_TO_GRAPHQL as any)[key](value);
  }
  return obj as Partial<IndexingRuleAttributes>;
};
const INDEXING_RULE_CONVERTERS_FROM_GRAPHQL: Record<
  keyof IndexingRuleAttributes,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  (x: never) => any
> = {
  id: (x) => x,
  identifier: (x) => x,
  identifierType: (x) => x,
  allocationAmount: nullPassThrough((x: string) => BigNumber.from(x)),
  allocationLifetime: nullPassThrough((x: string) => parseInt(x)),
  autoRenewal: (x) => x,
  parallelAllocations: nullPassThrough((x: string) => parseInt(x)),
  minSignal: nullPassThrough((x: string) => BigNumber.from(x)),
  maxSignal: nullPassThrough((x: string) => BigNumber.from(x)),
  minStake: nullPassThrough((x: string) => BigNumber.from(x)),
  maxAllocationPercentage: nullPassThrough((x: string) => parseFloat(x)),
  minAverageQueryFees: nullPassThrough((x: string) => BigNumber.from(x)),
  decisionBasis: (x) => x,
  custom: nullPassThrough(JSON.stringify),
  requireSupported: (x) => x,
};
/**
 * Parses an indexing rule returned from the indexer management GraphQL
 * API into normalized form.
 */
export const indexingRuleFromGraphQL = (
  rule: Partial<IndexingRuleAttributes>,
): Partial<IndexingRuleAttributes> => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const obj = {} as any;
  for (const [key, value] of Object.entries(pickFields(rule, []))) {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    obj[key] = (INDEXING_RULE_CONVERTERS_FROM_GRAPHQL as any)[key](value);
  }
  return obj as Partial<IndexingRuleAttributes>;
};
export const setIndexingRule = async (
  client: IndexerManagementClient,
  rule: Partial<IndexingRuleAttributes>,
): Promise<Partial<IndexingRuleAttributes>> => {
  const result = await client
    .mutation(
      gql`
        mutation setIndexingRule($rule: IndexingRuleInput!) {
          setIndexingRule(rule: $rule) {
            identifier
            identifierType
            allocationAmount
            allocationLifetime
            autoRenewal
            parallelAllocations
            maxAllocationPercentage
            minSignal
            maxSignal
            minStake
            minAverageQueryFees
            custom
            decisionBasis
            requireSupported
          }
        }
      `,
      { rule: indexingRuleToGraphQL(rule) },
    )
    .toPromise();

  if (result.error) {
    throw result.error;
  }

  return indexingRuleFromGraphQL(result.data.setIndexingRule);
};

export const indexingRules = async (
  client: IndexerManagementClient,
  merged: boolean,
): Promise<Partial<IndexingRuleAttributes>[]> => {
  const result = await client
    .query(
      gql`
        query indexingRules($merged: Boolean!) {
          indexingRules(merged: $merged) {
            identifier
            identifierType
            allocationAmount
            allocationLifetime
            autoRenewal
            parallelAllocations
            maxAllocationPercentage
            minSignal
            maxSignal
            minStake
            minAverageQueryFees
            custom
            decisionBasis
            requireSupported
          }
        }
      `,
      { merged: !!merged },
    )
    .toPromise();

  if (result.error) {
    throw result.error;
  }

  return result.data.indexingRules.map(indexingRuleFromGraphQL);
};
