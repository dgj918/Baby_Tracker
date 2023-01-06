import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EatingInputComponent } from './eating-input.component';

describe('EatingInputComponent', () => {
  let component: EatingInputComponent;
  let fixture: ComponentFixture<EatingInputComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EatingInputComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EatingInputComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
