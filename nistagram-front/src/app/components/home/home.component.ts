import { Component, OnInit } from '@angular/core';
import {PostModel} from '../../models/post.model';
import {Post} from '../profile/profile.component';

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
    this.post = {
      username: 'slavBrain',
      message: 'message1',
      likes: 10,
      comments: 20,
      date: new Date()
    };
    this.posts.push(this.post);
    this.posts.push(this.post);
    this.posts.push(this.post);
    this.posts.push(this.post);
  }

}
