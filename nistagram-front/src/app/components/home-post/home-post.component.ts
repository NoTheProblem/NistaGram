import { Component, OnInit, Input } from '@angular/core';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-home-post',
  templateUrl: './home-post.component.html',
  styleUrls: ['./home-post.component.css']
})
export class HomePostComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  imageAlbumNumber: number;

  constructor() { }

  ngOnInit(): void {
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
}
