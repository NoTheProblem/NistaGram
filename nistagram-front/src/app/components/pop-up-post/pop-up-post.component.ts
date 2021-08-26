import {Component, Input, OnInit} from '@angular/core';
import {PostModel} from '../../models/post.model';
import {PostService} from '../../services/post.service';
@Component({
  selector: 'app-pop-up-post',
  templateUrl: './pop-up-post.component.html',
  styleUrls: ['./pop-up-post.component.css']
})
export class PopUpPostComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  imageAlbumNumber: number;
  isLiked: boolean;
  commentInput: string;
  isCommentInput = false;
  showComments = false;


  constructor(private postService: PostService) { }

  ngOnInit(): void {
    this.imageAlbumNumber = 0;
    this.image = 'data:image/jpg;base64,' + this.post.images[0].Image;

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
    this.postService.reportPost(this.post.id);
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


}
