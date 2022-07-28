import { CreatePostGQL, CreatePostMutation, GetPostsGQL, GetPostByIdGQL } from './../../generated/graphql';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { MutationResult } from 'apollo-angular';

@Injectable({
  providedIn: 'root'
})
export class PostsService {

  constructor(
    private createPost: CreatePostGQL,
    private getPosts: GetPostsGQL,
    private getPostById: GetPostByIdGQL,
  ) { }

  onGetPosts () {
    return this.getPosts.fetch()
  }

  onGetPostById(id: string) {
    return this.getPostById.fetch({postId: id})
  }


  onCreatePost(values: any): Observable<MutationResult<CreatePostMutation>> {
    const {content, tittle, url} = values
    const post = this.createPost.mutate({
      input:{
        user: "culon",
        content: {
          body: content,
          tittle,
          image: url
        }
      }
    })
    return post
  }
}
