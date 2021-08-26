import {Component, Input, OnInit} from '@angular/core';
import {PENALTY, ReportModel} from '../../models/report.model';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {PostService} from '../../services/post.service';

@Component({
  selector: 'app-report-admin',
  templateUrl: './report-admin.component.html',
  styleUrls: ['./report-admin.component.css']
})
export class ReportAdminComponent implements OnInit {
  @Input() reportedPost: ReportModel;

  constructor(private modalService: NgbModal, private postService: PostService) { }

  ngOnInit(): void {
  }


  open(content): void {
    this.modalService.open(content,
      {centered: true, scrollable: true, size: 'xl'});
  }

  declineReport(): void {
    this.postService.answerReport({id: this.reportedPost.id, penalty: PENALTY.DECLINE_REPORT});
  }

  deleteAccount(): void {
    this.postService.answerReport({id: this.reportedPost.id, penalty: PENALTY.DELETE_PROFILE});
  }

  deletePost(): void {
    this.postService.answerReport({id: this.reportedPost.id, penalty: PENALTY.REMOVE_CONTENT});
  }
}
