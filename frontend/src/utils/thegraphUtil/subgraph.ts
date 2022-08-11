import { createClient, gql } from '@urql/core';

export function createSubgraphClient(url: string, options: Recordable = {}) {
  return createClient({ url, ...options });
}

// Subgraph API
export function fetchSubgraphs(page = 1, perPage = 20) {
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

  // const apiURL = 'https://gateway.thegraph.com/network'
  const apiURL = 'https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-testnet';
  const client = createSubgraphClient(apiURL);

  return client
    .query(tokensQuery, {
      first: perPage,
      skip: (page - 1) * perPage,
      orderBy: 'currentSignalledTokens',
      orderDirection: 'desc',
    })
    .toPromise();
}
