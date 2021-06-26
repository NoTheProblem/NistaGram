import { Component, OnInit } from '@angular/core';
import {newArray} from '@angular/compiler/src/util';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public post: PostModel;
  public posts: Array<PostModel> = new Array<PostModel>();
  constructor() { }

  ngOnInit(): void {
    this.post = new PostModel('slavBrain', 'message1', 10, 20);
    this.posts.push(this.post);
    this.post = new PostModel( 'slavBrain', 'message1', 10, 20);
    this.posts.push(this.post);
    this.post = new PostModel('slavBrain', 'message1', 10, 20);
    this.posts.push(this.post);
    this.post = new PostModel('slavBrain', 'message1', 10, 20);
    this.posts.push(this.post);
    this.post = new PostModel('slavBrain', 'message1', 10, 20);
    this.posts.push(this.post);
  }

}

export class PostModel{
  constructor(
    public username: string,
    public message: string,
    public likes: number,
    public comments: number,
    public date: Date = new Date()
  ) {
  }
}
