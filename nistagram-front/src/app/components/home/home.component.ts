import { Component, OnInit } from '@angular/core';
import {PostModel} from '../../models/post.model';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  public posts: Array<PostModel> = new Array<PostModel>();



  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.initPosts();
  }
  private initPosts(): void {
    this.postService.HomeFeed()
      .subscribe((postsList: Array<PostModel>) => {
        this.posts = postsList;
      });
  }
}



