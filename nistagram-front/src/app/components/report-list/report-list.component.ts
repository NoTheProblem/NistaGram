import { Component, OnInit } from '@angular/core';
import {TokenStorageService} from '../../_services/token-storage.service';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {PostService} from '../../services/post.service';
import {ReportModel} from '../../models/report.model';

@Component({
  selector: 'app-report-list',
  templateUrl: './report-list.component.html',
  styleUrls: ['./report-list.component.css']
})
export class ReportListComponent implements OnInit {

  reports: Array<ReportModel> = new Array<ReportModel>();
  i: number;

  constructor(private tokenStorageService: TokenStorageService, private modalService: NgbModal,
              private postService: PostService) { }

  ngOnInit(): void {
    if (this.tokenStorageService.getRole() !== 1){
      // TODO redirect
    }
    this.i = 0;
    this.postService.getUnAnsweredReports()
      .subscribe((reportList: Array<ReportModel>) => {
        this.reports = reportList;
      });
  }

  open(content, i): void {
    this.i = i;
    this.modalService.open(content,{centered: true, scrollable: true, size: 'xl'});
  }

}
