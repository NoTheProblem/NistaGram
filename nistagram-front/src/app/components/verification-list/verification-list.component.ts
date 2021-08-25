import { Component, OnInit } from '@angular/core';
import {TokenStorageService} from '../../_services/token-storage.service';
import {NgbModal} from '@ng-bootstrap/ng-bootstrap';
import {VerificationRequestModel} from '../../models/verification-request.model';
import {RequestService} from '../../services/request.service';
import {PostModel} from '../../models/post.model';

@Component({
  selector: 'app-verification-list',
  templateUrl: './verification-list.component.html',
  styleUrls: ['./verification-list.component.css']
})
export class VerificationListComponent implements OnInit {

  verificationRequests: Array<VerificationRequestModel> = new Array<VerificationRequestModel>();
  i: number;

  constructor(private tokenStorageService: TokenStorageService, private modalService: NgbModal,
              private requestService: RequestService) { }

  ngOnInit(): void {
    if (this.tokenStorageService.getRole() !== 1){
      // TODO redirect
    }
    this.i = 0;
    this.requestService.getUnAnsweredVerificationRequests()
      .subscribe((requestList: Array<VerificationRequestModel>) => {
        this.verificationRequests = requestList;
        console.log(requestList);
      });
  }

  open(content, i): void {
    this.i = i;
    this.modalService.open(content,
      {centered: true, scrollable: true, size: 'xl'});
  }

}
