import { Component, OnInit, Input } from '@angular/core';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {
  @Input() post: PostModel;

  constructor() { }

  ngOnInit(): void {
  }

}
