import { gql } from 'apollo-angular';
import { Injectable } from '@angular/core';
import * as Apollo from 'apollo-angular';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type AuthResponse = {
  __typename?: 'AuthResponse';
  token: Scalars['String'];
  user: User;
};

export type Content = {
  __typename?: 'Content';
  body: Scalars['String'];
  image: Scalars['String'];
  tittle: Scalars['String'];
};

export type LoginInput = {
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createPost: Post;
  deletePost: Post;
  deleteUser: User;
  login: AuthResponse;
  signIn: AuthResponse;
  updatePost: Post;
};


export type MutationCreatePostArgs = {
  input: NewPost;
};


export type MutationDeletePostArgs = {
  id: Scalars['String'];
};


export type MutationDeleteUserArgs = {
  id: Scalars['String'];
};


export type MutationLoginArgs = {
  input?: InputMaybe<LoginInput>;
};


export type MutationSignInArgs = {
  input?: InputMaybe<NewUser>;
};


export type MutationUpdatePostArgs = {
  input: UpdatePost;
};

export type NewContent = {
  body: Scalars['String'];
  image: Scalars['String'];
  tittle: Scalars['String'];
};

export type NewPost = {
  content: NewContent;
  user: Scalars['String'];
};

export type NewUser = {
  password: Scalars['String'];
  username: Scalars['String'];
};

export type Post = {
  __typename?: 'Post';
  content: Content;
  date: Scalars['String'];
  id: Scalars['String'];
  user: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  post: Post;
  posts: Array<Post>;
  user: User;
};


export type QueryPostArgs = {
  id: Scalars['String'];
};


export type QueryUserArgs = {
  id: Scalars['String'];
};

export type UpdatePost = {
  content: NewContent;
  id: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['String'];
  username: Scalars['String'];
};

export type SigninMutationVariables = Exact<{
  input?: InputMaybe<NewUser>;
}>;


export type SigninMutation = { __typename?: 'Mutation', signIn: { __typename?: 'AuthResponse', token: string, user: { __typename?: 'User', id: string, username: string } } };

export type LoginMutationVariables = Exact<{
  input?: InputMaybe<LoginInput>;
}>;


export type LoginMutation = { __typename?: 'Mutation', login: { __typename?: 'AuthResponse', token: string, user: { __typename?: 'User', id: string, username: string } } };

export type DeleteUserMutationVariables = Exact<{
  deleteUserId: Scalars['String'];
}>;


export type DeleteUserMutation = { __typename?: 'Mutation', deleteUser: { __typename?: 'User', id: string, username: string } };

export type CreatePostMutationVariables = Exact<{
  input: NewPost;
}>;


export type CreatePostMutation = { __typename?: 'Mutation', createPost: { __typename?: 'Post', id: string, user: string, date: string, content: { __typename?: 'Content', tittle: string, body: string, image: string } } };

export type UpdatePostMutationVariables = Exact<{
  input: UpdatePost;
}>;


export type UpdatePostMutation = { __typename?: 'Mutation', updatePost: { __typename?: 'Post', id: string, user: string, date: string, content: { __typename?: 'Content', tittle: string, body: string, image: string } } };

export type DeletePostMutationVariables = Exact<{
  deletePostId: Scalars['String'];
}>;


export type DeletePostMutation = { __typename?: 'Mutation', deletePost: { __typename?: 'Post', id: string, user: string, date: string, content: { __typename?: 'Content', tittle: string, body: string, image: string } } };

export type GetPostsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetPostsQuery = { __typename?: 'Query', posts: Array<{ __typename?: 'Post', id: string, user: string, date: string, content: { __typename?: 'Content', tittle: string, body: string, image: string } }> };

export type GetPostByIdQueryVariables = Exact<{
  postId: Scalars['String'];
}>;


export type GetPostByIdQuery = { __typename?: 'Query', post: { __typename?: 'Post', id: string, user: string, date: string, content: { __typename?: 'Content', tittle: string, body: string, image: string } } };

export type GetUserByIdQueryVariables = Exact<{
  userId: Scalars['String'];
}>;


export type GetUserByIdQuery = { __typename?: 'Query', user: { __typename?: 'User', id: string, username: string } };

export const SigninDocument = gql`
    mutation Signin($input: NewUser) {
  signIn(input: $input) {
    token
    user {
      id
      username
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class SigninGQL extends Apollo.Mutation<SigninMutation, SigninMutationVariables> {
            // @ts-ignore
    document = SigninDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const LoginDocument = gql`
    mutation Login($input: LoginInput) {
  login(input: $input) {
    token
    user {
      id
      username
    }
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class LoginGQL extends Apollo.Mutation<LoginMutation, LoginMutationVariables> {
            // @ts-ignore
    document = LoginDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeleteUserDocument = gql`
    mutation DeleteUser($deleteUserId: String!) {
  deleteUser(id: $deleteUserId) {
    id
    username
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DeleteUserGQL extends Apollo.Mutation<DeleteUserMutation, DeleteUserMutationVariables> {
            // @ts-ignore
    document = DeleteUserDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const CreatePostDocument = gql`
    mutation CreatePost($input: NewPost!) {
  createPost(input: $input) {
    id
    user
    content {
      tittle
      body
      image
    }
    date
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class CreatePostGQL extends Apollo.Mutation<CreatePostMutation, CreatePostMutationVariables> {
            // @ts-ignore
    document = CreatePostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const UpdatePostDocument = gql`
    mutation UpdatePost($input: UpdatePost!) {
  updatePost(input: $input) {
    id
    user
    content {
      tittle
      body
      image
    }
    date
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class UpdatePostGQL extends Apollo.Mutation<UpdatePostMutation, UpdatePostMutationVariables> {
            // @ts-ignore
    document = UpdatePostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const DeletePostDocument = gql`
    mutation DeletePost($deletePostId: String!) {
  deletePost(id: $deletePostId) {
    id
    user
    content {
      tittle
      body
      image
    }
    date
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class DeletePostGQL extends Apollo.Mutation<DeletePostMutation, DeletePostMutationVariables> {
            // @ts-ignore
    document = DeletePostDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const GetPostsDocument = gql`
    query GetPosts {
  posts {
    id
    user
    content {
      tittle
      body
      image
    }
    date
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class GetPostsGQL extends Apollo.Query<GetPostsQuery, GetPostsQueryVariables> {
            // @ts-ignore
    document = GetPostsDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const GetPostByIdDocument = gql`
    query GetPostById($postId: String!) {
  post(id: $postId) {
    id
    user
    content {
      tittle
      body
      image
    }
    date
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class GetPostByIdGQL extends Apollo.Query<GetPostByIdQuery, GetPostByIdQueryVariables> {
            // @ts-ignore
    document = GetPostByIdDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }
export const GetUserByIdDocument = gql`
    query GetUserById($userId: String!) {
  user(id: $userId) {
    id
    username
  }
}
    `;

  @Injectable({
    providedIn: 'root'
  })
  export class GetUserByIdGQL extends Apollo.Query<GetUserByIdQuery, GetUserByIdQueryVariables> {
            // @ts-ignore
    document = GetUserByIdDocument;
    
    constructor(apollo: Apollo.Apollo) {
      super(apollo);
    }
  }