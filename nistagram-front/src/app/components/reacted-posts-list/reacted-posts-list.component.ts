import { Component, OnInit } from '@angular/core';
import {PostService} from '../../services/post.service';
import {PostReactionsModel} from '../../models/post-reactions.model';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-reacted-posts-list',
  templateUrl: './reacted-posts-list.component.html',
  styleUrls: ['./reacted-posts-list.component.css']
})
export class ReactedPostsListComponent implements OnInit {
  reactedPosts: PostReactionsModel;
  post: PostModel;
  i: number;

  constructor(private postService: PostService, private modalService: NgbModal) { }

  ngOnInit(): void {
    this.postService.getReactedPosts().subscribe((reactions: PostReactionsModel) => {
      this.reactedPosts = reactions;
    }, error => {
      console.log(error);
    });
  }

  open(content, i: number, b: boolean): void {
    this.i = i;
    if (b){
      this.post = this.reactedPosts.likedPosts[i];
    }
    else{
      this.post = this.reactedPosts.dislikedPosts[i];
    }
    this.modalService.open(content,
      {centered: true, scrollable: true, size: 'xl'});
  }

}
