import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CurrentFeedingStatusComponent } from './current-feeding-status.component';

describe('CurrentFeedingStatusComponent', () => {
  let component: CurrentFeedingStatusComponent;
  let fixture: ComponentFixture<CurrentFeedingStatusComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CurrentFeedingStatusComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CurrentFeedingStatusComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
