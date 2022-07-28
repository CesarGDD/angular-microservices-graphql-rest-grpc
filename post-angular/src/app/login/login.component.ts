import { NgForm } from '@angular/forms';
import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { User } from 'src/generated/graphql';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
})
export class LoginComponent implements OnInit {
  constructor(
    private authSvc: AuthService,
    ) {}

  newUser = false

  onSwichMode() {
    this.newUser = !this.newUser
  }

  auth(form: NgForm) {
    const { username, password } = form.value;
    if (this.newUser){
      this.authSvc.onSignin(username, password)
    }else {
      this.authSvc.onLogin(username, password)
    }
  }

  ngOnInit(): void {}
}
