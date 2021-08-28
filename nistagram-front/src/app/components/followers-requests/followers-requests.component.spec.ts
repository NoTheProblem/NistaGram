import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FollowersRequestsComponent } from './followers-requests.component';

describe('FollowersRequestsComponent', () => {
  let component: FollowersRequestsComponent;
  let fixture: ComponentFixture<FollowersRequestsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FollowersRequestsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FollowersRequestsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
