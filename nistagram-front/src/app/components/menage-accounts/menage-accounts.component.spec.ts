import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenageAccountsComponent } from './menage-accounts.component';

describe('MenageAccountsComponent', () => {
  let component: MenageAccountsComponent;
  let fixture: ComponentFixture<MenageAccountsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MenageAccountsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MenageAccountsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
