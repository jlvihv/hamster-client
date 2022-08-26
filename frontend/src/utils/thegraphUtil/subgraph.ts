import { createClient, gql, Client } from '@urql/core';
// import { parseIndexingRule, setIndexingRule, indexingRules } from './rules';
// import {
//   processIdentifier,
//   IndexingDecisionBasis,
//   IndexerManagementClient,
// } from '@graphprotocol/indexer-common';

export function createSubgraphClient(url: string, options: Recordable = {}) {
  return createClient({ url, ...options });
}

// Subgraph API
export function fetchSubgraphs(client: Client, page = 1, perPage = 20) {
  const tokensQuery = gql`
    query subgraphs(
      $orderBy: Subgraph_orderBy
      $orderDirection: OrderDirection
      $first: Int
      $skip: Int
    ) {
      subgraphs(
        orderBy: $orderBy
        orderDirection: $orderDirection
        first: $first
        skip: $skip
        where: { active: true, displayName_not: "", entityVersion: 2 }
      ) {
        id
        entityVersion
        displayName
        description
        image
        createdAt
        updatedAt
        signalledTokens
        currentSignalledTokens
        active
        owner {
          id
          image
          defaultName {
            id
            name
            __typename
          }
          __typename
        }
        currentVersion {
          id
          subgraphDeployment {
            id
            ipfsHash
            stakedTokens
            signalledTokens
            queryFeesAmount
            __typename
          }
          __typename
        }
        pastVersions {
          id
          createdAt
          label
          subgraphDeployment {
            id
            signalledTokens
            __typename
          }
          __typename
        }
        __typename
      }
    }
  `;

  return client
    .query(tokensQuery, {
      first: perPage,
      skip: (page - 1) * perPage,
      orderBy: 'currentSignalledTokens',
      orderDirection: 'desc',
    })
    .toPromise();
}

// export async function deploySubgraph(client: Client, deploymentId: string) {
//   const [identifier, identifierType] = await processIdentifier(deploymentId, {
//     all: false,
//     global: true,
//   });

//   const inputRule = parseIndexingRule({
//     identifier,
//     identifierType,
//     decisionBasis: IndexingDecisionBasis.ALWAYS,
//   });

//   await setIndexingRule(client as IndexerManagementClient, inputRule);
// }

// export function getDeployedSubgraphs(client: Client) {
//   return indexingRules(client as IndexerManagementClient, false);
// }
