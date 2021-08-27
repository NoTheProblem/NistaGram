import { Component, OnInit } from '@angular/core';
import {PostService} from '../../services/post.service';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-home-explore',
  templateUrl: './home-explore.component.html',
  styleUrls: ['./home-explore.component.css']
})
export class HomeExploreComponent implements OnInit {
  public posts: Array<PostModel> = new Array<PostModel>();

  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.postService.explore()
      .subscribe((postsList: Array<PostModel>) => {
        this.posts = postsList;
      });
  }

}
