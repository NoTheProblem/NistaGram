import {Component, Input, OnInit} from '@angular/core';
import {PostModel} from '../../models/post.model';
import {UserModel} from '../../models/user.model';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-privacy-settings',
  templateUrl: './privacy-settings.component.html',
  styleUrls: ['./privacy-settings.component.css']
})
export class PrivacySettingsComponent implements OnInit {
  @Input() userProfile: UserModel;
  privacy: boolean;
  messages: boolean;
  tag: boolean;

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {
    this.privacy = this.userProfile.profilePrivacy;
    this.messages = this.userProfile.receiveMessages;
    this.tag = this.userProfile.taggable;
  }

  updateProfile(): void {
    this.settingsService.updatePrivacySettings(this.privacy, this.messages, this.tag);

  }
}
