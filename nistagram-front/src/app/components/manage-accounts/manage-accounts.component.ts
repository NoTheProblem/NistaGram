import { Component, OnInit } from '@angular/core';
import {AuthService} from '../../_services/auth.service';
import {BusinessRequestModel} from '../../models/business-request.model';
import {STATUS_TYPE} from '../../models/business-request.model';
import {USER_TYPE, UserModel} from '../../models/user.model';

@Component({
  selector: 'app-menage-accounts',
  templateUrl: './manage-accounts.component.html',
  styleUrls: ['./manage-accounts.component.css']
})
export class ManageAccountsComponent implements OnInit {

  businessRequests: Array<BusinessRequestModel> = new Array<BusinessRequestModel>();
  users: Array<UserModel> = new Array<UserModel>();
  isBusinessRequests = false;
  isUsers = false;
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

  getAllUsers(): void {
    this.isUsers = !this.isUsers;
    if (this.isUsers){
      this.authService.getAllUsers().subscribe((users: Array<UserModel>) => {
        this.users = users;
      }, error => {
        console.log(error);
      });
    }
    else{
      this.users = new Array<UserModel>();
    }
  }

  makeRole(i: number, roleNumber: number): void {
    const answer = {
      username: this.users[i].username,
      role: roleNumber
    };
    this.users[i].role = roleNumber;
    this.authService.chagneRole(answer);

  }
}
