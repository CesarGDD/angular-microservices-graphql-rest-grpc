import { CreatePostGQL, GetPostsGQL } from './../../generated/graphql';
import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { PostsService } from '../services/posts.service';

@Component({
  selector: 'app-newpost',
  templateUrl: './newpost.component.html',
  styleUrls: ['./newpost.component.css']
})
export class NewpostComponent implements OnInit {

  constructor(
    private getPosts: GetPostsGQL,
    private postSvc: PostsService,
  ) { }



  onCreatePost(form: NgForm){
  this.postSvc.onCreatePost(form.value).subscribe(({data, loading, errors}) => {
      console.log(data);
      form.reset()
    })

  }

  ngOnInit(): void {
    this.getPosts.fetch().subscribe(({data, errors}) => {      
      console.log(data);
      
    })
  }

}
