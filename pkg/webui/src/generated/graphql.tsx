import * as React from 'react';
import * as Apollo from '@apollo/client';
import * as ApolloReactComponents from '@apollo/client/react/components';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
const gql = Apollo.gql;
export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  int: any;
  string: any;
};

export type Asset = {
  __typename?: 'Asset';
  /** @deprecated  */
  copyright: Copyright;
  /** @deprecated  */
  extracts: Extracts;
  /** @deprecated  */
  filename: Scalars['string'];
  /** @deprecated  */
  genre: Scalars['string'];
  /** @deprecated  */
  id: Scalars['string'];
  /** @deprecated  */
  media: Media;
  /** @deprecated  */
  path: Scalars['string'];
  /** @deprecated  */
  ratings: Ratings;
  /** @deprecated  */
  referenceCopies: Array<Scalars['string']>;
  /** @deprecated  */
  title: Scalars['string'];
};

export type Copyright = {
  __typename?: 'Copyright';
  /** @deprecated  */
  lyrics: Scalars['string'];
  /** @deprecated  */
  music: Scalars['string'];
};

export type Extracts = {
  __typename?: 'Extracts';
  /** @deprecated  */
  available: Array<Scalars['int']>;
  /** @deprecated  */
  preselected: Array<Scalars['int']>;
};

export type Media = {
  __typename?: 'Media';
  /** @deprecated  */
  self: Scalars['string'];
  /** @deprecated  */
  thumbnail: Scalars['string'];
};

export type Project = {
  __typename?: 'Project';
  /** @deprecated  */
  assets: Array<ProjectAsset>;
  /** @deprecated  */
  productionNode: Scalars['string'];
  /** @deprecated  */
  title: Scalars['string'];
};

export type ProjectAsset = {
  __typename?: 'ProjectAsset';
  /** @deprecated  */
  asset: Asset;
  /** @deprecated  */
  rating: Scalars['int'];
  /** @deprecated  */
  referenceCopies: Array<Scalars['string']>;
  /** @deprecated  */
  selectedExtracts: Array<Scalars['int']>;
  /** @deprecated  */
  sortString: Scalars['string'];
};

export type Query = {
  __typename?: 'Query';
  /** @deprecated  */
  asset?: Maybe<Asset>;
  /** @deprecated  */
  assets: Array<Asset>;
  /** @deprecated  */
  projects: Array<Project>;
};


export type QueryAssetArgs = {
  filename: Scalars['string'];
};

export type Ratings = {
  __typename?: 'Ratings';
  /** @deprecated  */
  difficulty: Scalars['int'];
  /** @deprecated  */
  ensemblePlayability: Scalars['int'];
};



export type AssetQueryVariables = Exact<{
  filename: Scalars['string'];
}>;


export type AssetQuery = (
  { __typename?: 'Query' }
  & { asset?: Maybe<(
    { __typename?: 'Asset' }
    & Pick<Asset, 'id' | 'title' | 'filename' | 'path'>
    & { media: (
      { __typename?: 'Media' }
      & Pick<Media, 'self'>
    ) }
  )> }
);

export type AssetsQueryVariables = Exact<{ [key: string]: never; }>;


export type AssetsQuery = (
  { __typename?: 'Query' }
  & { assets: Array<(
    { __typename?: 'Asset' }
    & Pick<Asset, 'id' | 'title' | 'filename' | 'path'>
  )> }
);


export const AssetDocument = gql`
    query Asset($filename: string!) {
  asset(filename: $filename) {
    id
    title
    filename
    path
    media {
      self
    }
  }
}
    `;
export type AssetComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<AssetQuery, AssetQueryVariables>, 'query'> & ({ variables: AssetQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const AssetComponent = (props: AssetComponentProps) => (
      <ApolloReactComponents.Query<AssetQuery, AssetQueryVariables> query={AssetDocument} {...props} />
    );
    

/**
 * __useAssetQuery__
 *
 * To run a query within a React component, call `useAssetQuery` and pass it any options that fit your needs.
 * When your component renders, `useAssetQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAssetQuery({
 *   variables: {
 *      filename: // value for 'filename'
 *   },
 * });
 */
export function useAssetQuery(baseOptions?: Apollo.QueryHookOptions<AssetQuery, AssetQueryVariables>) {
        return Apollo.useQuery<AssetQuery, AssetQueryVariables>(AssetDocument, baseOptions);
      }
export function useAssetLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AssetQuery, AssetQueryVariables>) {
          return Apollo.useLazyQuery<AssetQuery, AssetQueryVariables>(AssetDocument, baseOptions);
        }
export type AssetQueryHookResult = ReturnType<typeof useAssetQuery>;
export type AssetLazyQueryHookResult = ReturnType<typeof useAssetLazyQuery>;
export type AssetQueryResult = Apollo.QueryResult<AssetQuery, AssetQueryVariables>;
export const AssetsDocument = gql`
    query Assets {
  assets {
    id
    title
    filename
    path
  }
}
    `;
export type AssetsComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<AssetsQuery, AssetsQueryVariables>, 'query'>;

    export const AssetsComponent = (props: AssetsComponentProps) => (
      <ApolloReactComponents.Query<AssetsQuery, AssetsQueryVariables> query={AssetsDocument} {...props} />
    );
    

/**
 * __useAssetsQuery__
 *
 * To run a query within a React component, call `useAssetsQuery` and pass it any options that fit your needs.
 * When your component renders, `useAssetsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAssetsQuery({
 *   variables: {
 *   },
 * });
 */
export function useAssetsQuery(baseOptions?: Apollo.QueryHookOptions<AssetsQuery, AssetsQueryVariables>) {
        return Apollo.useQuery<AssetsQuery, AssetsQueryVariables>(AssetsDocument, baseOptions);
      }
export function useAssetsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AssetsQuery, AssetsQueryVariables>) {
          return Apollo.useLazyQuery<AssetsQuery, AssetsQueryVariables>(AssetsDocument, baseOptions);
        }
export type AssetsQueryHookResult = ReturnType<typeof useAssetsQuery>;
export type AssetsLazyQueryHookResult = ReturnType<typeof useAssetsLazyQuery>;
export type AssetsQueryResult = Apollo.QueryResult<AssetsQuery, AssetsQueryVariables>;