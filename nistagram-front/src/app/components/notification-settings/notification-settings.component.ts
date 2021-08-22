import {Component, Input, OnInit} from '@angular/core';
import {UserModel} from '../../models/user.model';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-notification-settings',
  templateUrl: './notification-settings.component.html',
  styleUrls: ['./notification-settings.component.css']
})
export class NotificationSettingsComponent implements OnInit {
  @Input() userProfile: UserModel;
  commentNotifications: boolean;
  messageNotifications: boolean;
  postNotifications: boolean;

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {
    this.commentNotifications =  this.userProfile.receiveCommentNotifications;
    this.messageNotifications = this.userProfile.receiveMessagesNotifications;
    this.postNotifications = this.userProfile.receivePostNotifications;

  }

  updateProfile(): void {
    this.settingsService.updateNotificationSettings(this.commentNotifications, this.messageNotifications, this.postNotifications);
  }
}
