import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import {AuthInterceptor} from './_helpers/auth.interceptor';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import {MatCardModule} from '@angular/material/card';
import {FormsModule} from '@angular/forms';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatDialogModule} from '@angular/material/dialog';
import { NavigationComponent } from './components/navigation/navigation.component';
import {MatIconModule} from '@angular/material/icon';
import {MatMenuModule} from '@angular/material/menu';
import {MatBadgeModule} from '@angular/material/badge';
import { ProfileComponent } from './components/profile/profile.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { HomePostComponent } from './components/home-post/home-post.component';
import { UploadPostComponent } from './components/upload-post/upload-post.component';
import { ToastrModule } from 'ngx-toastr';
import { SettingsComponent } from './components/settings/settings.component';
import {MatSidenavModule} from '@angular/material/sidenav';
import { EditProfileComponent } from './components/edit-profile/edit-profile.component';
import {MatDividerModule} from '@angular/material/divider';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatNativeDateModule} from '@angular/material/core';
import { NotificationSettingsComponent } from './components/notification-settings/notification-settings.component';
import { PrivacySettingsComponent } from './components/privacy-settings/privacy-settings.component';
import { ResetPasswordComponent } from './components/reset-password/reset-password.component';
import { VerificationComponent } from './components/verification/verification.component';
import { ProfilePostComponent } from './components/profile-post/profile-post.component';
import { PopUpPostComponent } from './components/pop-up-post/pop-up-post.component';
import { VerificationAdminComponent } from './components/verification-admin/verification-admin.component';
import { VerificationListComponent } from './components/verification-list/verification-list.component';
import { ReportListComponent } from './components/report-list/report-list.component';
import { ReportAdminComponent } from './components/report-admin/report-admin.component';
import { SearchListComponent } from './components/search-list/search-list.component';
import { SuggestedComponent } from './components/suggested/suggested.component';
import { ReactedPostsListComponent } from './components/reacted-posts-list/reacted-posts-list.component';
import { HomeExploreComponent } from './components/home-explore/home-explore.component';
import { ManageAccountsComponent } from './components/manage-accounts/manage-accounts.component';
import { FollowersRequestsComponent } from './components/followers-requests/followers-requests.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent,
    NavigationComponent,
    ProfileComponent,
    HomePostComponent,
    UploadPostComponent,
    SettingsComponent,
    EditProfileComponent,
    NotificationSettingsComponent,
    PrivacySettingsComponent,
    ResetPasswordComponent,
    VerificationComponent,
    ProfilePostComponent,
    PopUpPostComponent,
    VerificationAdminComponent,
    VerificationListComponent,
    ReportListComponent,
    ReportAdminComponent,
    SearchListComponent,
    SuggestedComponent,
    ReactedPostsListComponent,
    HomeExploreComponent,
    ManageAccountsComponent,
    FollowersRequestsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatCardModule,
    FormsModule,
    MatFormFieldModule,
    MatToolbarModule,
    MatDialogModule,
    MatIconModule,
    MatMenuModule,
    MatBadgeModule,
    NgbModule,
    FontAwesomeModule,
    HttpClientModule,
    ToastrModule.forRoot(),
    MatSidenavModule,
    MatDividerModule,
    MatDatepickerModule,
    MatNativeDateModule
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
