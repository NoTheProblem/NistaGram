import {Component, Input, OnInit} from '@angular/core';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-profile-picture',
  templateUrl: './profile-picture.component.html',
  styleUrls: ['./profile-picture.component.css']
})
export class ProfilePictureComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  constructor() { }

  ngOnInit(): void {
    if (this.post.comments?.length === undefined){
      this.post.numberOfComments = 0;
    }else {
      this.post.numberOfComments = this.post.comments.length;
    }
    this.image = 'data:image/jpg;base64,' + this.post.image;
  }

}


