import { Component, OnInit, Input } from '@angular/core';
import {PostModel} from '../../models/post.model';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-home-post',
  templateUrl: './home-post.component.html',
  styleUrls: ['./home-post.component.css']
})
export class HomePostComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  imageAlbumNumber: number;
  isLiked: boolean;
  commentInput: string;
  isCommentInput = false;
  showComments = false;

  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.isLiked = false;
    // TODO check if you already liked it
    if (this.post.comments?.length === undefined){
      this.post.numberOfComments = 0;
    }else {
      this.post.numberOfComments = this.post.comments.length;
    }
    this.imageAlbumNumber = 0;
    this.image = 'data:image/jpg;base64,' + this.post.images[0].Image;

  }

  albumLeft(): void {
    if (this.post.isAlbum){
      if (this.imageAlbumNumber !== 0){
        this.imageAlbumNumber = this.imageAlbumNumber - 1;
        this.image = 'data:image/jpg;base64,' + this.post.images[this.imageAlbumNumber].Image;
      }
    }
  }

  albumRight(): void {
    if (this.post.isAlbum){
      if (this.imageAlbumNumber !== this.post.images.length - 1){
        this.imageAlbumNumber = this.imageAlbumNumber + 1;
        this.image = 'data:image/jpg;base64,' + this.post.images[this.imageAlbumNumber].Image;
      }
    }

  }

  likePost(): void {
    this.isLiked = true;
    this.post.NumberOfLikes = this.post.NumberOfLikes + 1;
    this.postService.likePost({id: this.post.id});
  }

  unLikePost(): void {
    this.isLiked = false;
    this.post.NumberOfLikes = this.post.NumberOfLikes - 1;
    this.postService.disLikePost({id: this.post.id});
  }

  addComment(): void {
    this.isCommentInput = false;
    this.postService.commentPost({id: this.post.id, text: this.commentInput, date: new Date()});
    this.commentInput = null;

  }

  reportPost(): void {
    alert('Post reported');
    // TODO poziv backa
  }
}
