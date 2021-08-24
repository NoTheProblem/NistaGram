import { Component, OnInit, Input } from '@angular/core';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
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
