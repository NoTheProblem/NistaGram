import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReactedPostsListComponent } from './reacted-posts-list.component';

describe('ReactedPostsListComponent', () => {
  let component: ReactedPostsListComponent;
  let fixture: ComponentFixture<ReactedPostsListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReactedPostsListComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ReactedPostsListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
