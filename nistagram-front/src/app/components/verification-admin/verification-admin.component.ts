import {Component, Input, OnInit} from '@angular/core';
import {VerificationRequestModel} from '../../models/verification-request.model';
import {RequestService} from '../../services/request.service';

@Component({
  selector: 'app-verification-admin',
  templateUrl: './verification-admin.component.html',
  styleUrls: ['./verification-admin.component.css']
})
export class VerificationAdminComponent implements OnInit {
  @Input() request: VerificationRequestModel;
  image: any;
  openDeclineInput = false;
  declineInput: string;

  constructor(private requestService: RequestService) { }

  ngOnInit(): void {
    setTimeout(() => {
      this.image = 'data:image/jpg;base64,' + this.request.image;
    }, 500);
  }

  declineRequest(): void {
    this.requestService.sendVerificationAnswer({verificationAnswer: false, id: this.request.id, answer: this.declineInput});
  }

  acceptRequest(): void {
    this.requestService.sendVerificationAnswer({verificationAnswer: true, id: this.request.id, answer: ''});
  }
}
