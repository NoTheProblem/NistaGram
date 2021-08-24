import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PopUpPostComponent } from './pop-up-post.component';

describe('PopUpPostComponent', () => {
  let component: PopUpPostComponent;
  let fixture: ComponentFixture<PopUpPostComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PopUpPostComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PopUpPostComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
