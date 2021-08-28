import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './components/login/login.component';
import {RegisterComponent} from './components/register/register.component';
import {HomeComponent} from './components/home/home.component';
import {ProfileComponent} from './components/profile/profile.component';
import {UploadPostComponent} from './components/upload-post/upload-post.component';
import {SettingsComponent} from './components/settings/settings.component';
import {VerificationListComponent} from './components/verification-list/verification-list.component';
import {ReportListComponent} from './components/report-list/report-list.component';
import {SearchListComponent} from './components/search-list/search-list.component';
import {ReactedPostsListComponent} from './components/reacted-posts-list/reacted-posts-list.component';
import {ManageAccountsComponent} from './components/manage-accounts/manage-accounts.component';
import {FollowersRequestsComponent} from './components/followers-requests/followers-requests.component';

const routes: Routes = [

  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'home', component: HomeComponent},
  {path: 'upload-post', component: UploadPostComponent},
  {path: 'settings', component: SettingsComponent},
  {path: 'verification-list', component: VerificationListComponent},
  {path: 'report-list', component: ReportListComponent},
  {path: 'search', component: SearchListComponent},
  {path: 'reactions', component: ReactedPostsListComponent},
  {path: 'manage-accounts', component: ManageAccountsComponent},
  {path: 'followers-requests', component: FollowersRequestsComponent},

  // Ovaj mora da bude poslednji za sada
  {path: ':username', component: ProfileComponent},
  {path: '', redirectTo: 'home', pathMatch: 'full'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
