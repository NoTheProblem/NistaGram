import { Component, OnInit, Input } from '@angular/core';
import {PostModel} from '../../models/post.model';
import {PostService} from '../../services/post.service';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {TokenStorageService} from '../../_services/token-storage.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-home-post',
  templateUrl: './home-post.component.html',
  styleUrls: ['./home-post.component.css']
})
export class HomePostComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  imageAlbumNumber: number;
  isLiked = false;
  isDisLiked = false;
  commentInput: string;
  isCommentInput = false;
  showComments = false;

  constructor(private postService: PostService, private modalService: NgbModal,
              private tokenStorageService: TokenStorageService, private router: Router) { }

  ngOnInit(): void {
    if (this.tokenStorageService.isLoggedIn()){
      const username = this.tokenStorageService.getUsername();
      if (this.post.usersLiked !== null){
        if (this.post.usersLiked.includes(username)) {
          this.isLiked = true;
        }
      }
      if (this.post.usersDisliked !== null){
        if (this.post.usersDisliked.includes(username)){
          this.isDisLiked = true;
        }
      }
    }
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

  addComment(): void {
    this.isCommentInput = false;
    this.postService.commentPost({id: this.post.id, text: this.commentInput, date: new Date()});
    this.commentInput = null;

  }

  reportPost(): void {
    this.postService.reportPost(this.post.id);
  }

  open(content): void {
    this.modalService.open(content,
      {centered: true, scrollable: true, size: 'xl'});
  }

  likePost(): void {
    this.isLiked = true;
    this.post.NumberOfLikes = this.post.NumberOfLikes + 1;
    this.postService.likePost({id: this.post.id});
  }

  unLikePost(): void {
    this.isLiked = false;
    this.post.NumberOfLikes = this.post.NumberOfLikes - 1;
    this.postService.likePost({id: this.post.id});
  }

  disLikePost(): void {
    this.isDisLiked = true;
    this.post.NumberOfDislikes = this.post.NumberOfDislikes + 1;
    this.postService.disLikePost({id: this.post.id});
  }

  unDisLikePost(): void {
    this.isDisLiked = false;
    this.post.NumberOfDislikes = this.post.NumberOfDislikes - 1;
    this.postService.disLikePost({id: this.post.id});
  }

  search(s: string): void {
    this.router.navigate(['search'], { state: { searchQuery: s }});
  }
}
