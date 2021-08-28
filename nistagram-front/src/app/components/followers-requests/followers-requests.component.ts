import { Component, OnInit } from '@angular/core';
import {BusinessRequestModel, STATUS_TYPE} from '../../models/business-request.model';
import {FollowService} from '../../services/follow.service';
import {UsernameListModel} from '../../models/username-list.model';

@Component({
  selector: 'app-followers-requests',
  templateUrl: './followers-requests.component.html',
  styleUrls: ['./followers-requests.component.css']
})
export class FollowersRequestsComponent implements OnInit {
  requests: UsernameListModel;

  constructor(private followService: FollowService) { }

  ngOnInit(): void {
  }

  getRequests(): void {
    this.followService.getPendingRequests().subscribe((requests: UsernameListModel) => {
      this.requests = requests;
    }, error => {
      console.log(error);
    });
  }
  confirmRequest(i: number): void {
    this.followService.acceptRequest(this.requests.usernames[i]);

  }

  declineRequest(i: number): void {
    this.followService.unFollowUser(this.requests.usernames[i]);
  }

}

