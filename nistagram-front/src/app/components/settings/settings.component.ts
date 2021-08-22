import { Component, OnInit } from '@angular/core';
import {SettingsService} from '../../services/settings.service';
import {UserModel} from '../../models/user.model';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.css']
})
export class SettingsComponent implements OnInit {
  openEdit: boolean;
  openPassword: boolean;
  openPrivacy: boolean;
  openNotification: boolean;
  user: UserModel;

  constructor(private settingsService: SettingsService) { }

  ngOnInit(): void {
    this.settingsService.loadMyProfile()
      .subscribe((profile: UserModel) => {
        console.log(profile);
        this.user = profile;
        console.log(this.user);
      });
    this.openEdit = true;
    this.openNotification = false;
    this.openPrivacy = false;
    this.openPassword = false;
  }

  openTab(tab: string): void {
    switch (tab) {
      case 'edit' :
        this.openEdit = true;
        this.openNotification = false;
        this.openPrivacy = false;
        this.openPassword = false;
        break;
      case 'pw':
        this.openEdit = false;
        this.openNotification = false;
        this.openPrivacy = false;
        this.openPassword = true;
        break;

      case 'privacy' :
        this.openEdit = false;
        this.openNotification = false;
        this.openPrivacy = true;
        this.openPassword = false;
        break;

      case  'notification':
        this.openEdit = false;
        this.openNotification = true;
        this.openPrivacy = false;
        this.openPassword = false;
        break;

      default:
        this.openEdit = true;
        this.openNotification = false;
        this.openPrivacy = false;
        this.openPassword = false;
        break;
      }
    }
}
