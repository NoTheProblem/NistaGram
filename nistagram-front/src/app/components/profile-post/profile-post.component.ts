import {Component, Input, OnInit} from '@angular/core';
import {PostModel} from '../../models/post.model';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-profile-post',
  templateUrl: './profile-post.component.html',
  styleUrls: ['./profile-post.component.css']
})
export class ProfilePostComponent implements OnInit {
  @Input() post: PostModel;
  image: any;
  constructor(private modalService: NgbModal) { }

  ngOnInit(): void {
    if (this.post.comments?.length === undefined){
      this.post.numberOfComments = 0;
    }else {
      this.post.numberOfComments = this.post.comments.length;
    }
    this.image = 'data:image/jpg;base64,' + this.post.images[0].Image;
  }

  open(content): void {
    this.modalService.open(content,
      {centered: true, scrollable: true, size: 'xl'});
  }

}


