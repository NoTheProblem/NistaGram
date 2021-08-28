import { Component, OnInit } from '@angular/core';
import {AuthService} from '../../_services/auth.service';
import {BusinessRequestModel} from '../../models/business-request.model';
import {STATUS_TYPE} from '../../models/business-request.model';

@Component({
  selector: 'app-menage-accounts',
  templateUrl: './menage-accounts.component.html',
  styleUrls: ['./menage-accounts.component.css']
})
export class MenageAccountsComponent implements OnInit {

  businessRequests: Array<BusinessRequestModel> = new Array<BusinessRequestModel>();
  isBusinessRequests = false;
  constructor(private authService: AuthService) { }

  ngOnInit(): void {
  }

  getBusinessRequests(): void {
    this.isBusinessRequests = !this.isBusinessRequests;
    if (this.isBusinessRequests){
      this.authService.getPendingBusinessRequests().subscribe((requests: Array<BusinessRequestModel>) => {
        this.businessRequests = requests;
      }, error => {
        console.log(error);
      });
    }
    else{
      this.businessRequests = new Array<BusinessRequestModel>();
    }
  }

  confirmRequest(i: number): void {
    const answer = {
      username: this.businessRequests[i].username,
      status: STATUS_TYPE.ACCEPT
    };
    this.authService.answerBusinessRequest(answer);
    this.businessRequests = this.businessRequests.filter(request => request.username !== this.businessRequests[i].username);

  }

  declineRequest(i: number): void {
    const answer = {
      username: this.businessRequests[i].username,
      status: STATUS_TYPE.DECLINE
    };
    this.authService.answerBusinessRequest(answer);
    this.businessRequests = this.businessRequests.filter(request => request.username !== this.businessRequests[i].username);

  }

}
