import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SleepInputComponent } from './sleep-input.component';

describe('SleepInputComponent', () => {
  let component: SleepInputComponent;
  let fixture: ComponentFixture<SleepInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SleepInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SleepInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
