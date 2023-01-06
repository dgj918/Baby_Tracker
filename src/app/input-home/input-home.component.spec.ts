import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InputHomeComponent } from './input-home.component';

describe('InputHomeComponent', () => {
  let component: InputHomeComponent;
  let fixture: ComponentFixture<InputHomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InputHomeComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InputHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
