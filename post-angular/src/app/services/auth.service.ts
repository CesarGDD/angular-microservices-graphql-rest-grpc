import { LoginGQL, SigninGQL, User } from './../../generated/graphql';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';


@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(
    private login: LoginGQL,
    private signin: SigninGQL,
    private router: Router
  ) { }

  user: User = {
    id: '',
    username: '',
  };
  token: string | null = null;



  autoLogin() {
    const userData = localStorage.getItem('userData')
    const token = localStorage.getItem('token')
    if (userData && token) {
      const setUser:{id:string, username: string} = JSON.parse(userData)
      const setToken:string = JSON.parse(token)
      this.user = setUser
      this.token = setToken
      this.router.navigateByUrl('/')
    }
    if (!userData || !token){
      return
    }
  }

  onLogin (username: string, password: string): void {
    this.login.mutate({input: {
      username,
      password
    }}).subscribe({
      next: ({ data, loading }) => {
        console.log(data?.login);
        if (data) {
          this.user.id = data.login.user.id;
          this.user.username = data.login.user.username;
          this.token = data.login.token;
        }
      },
      error: (err) => {
        console.log("error message", err);
      },
      complete: () => {
        localStorage.setItem('userData', JSON.stringify(this.user))
        localStorage.setItem('token', JSON.stringify(this.token))
        this.router.navigateByUrl('/')
      }
    });
  }

  onSignin (username: string, password: string): void {
    this.signin.mutate( {
      input: {
        username,
        password
      }
    }).subscribe({
      next: ({ data, loading }) => {
        console.log(data?.signIn);
        if (data) {
          this.user.id = data.signIn.user.id;
          this.user.username = data.signIn.user.username;
          this.token = data.signIn.token;
        }
      },
      error: (err) => {
        console.log("error message", err);
      },
      complete: () => {
        localStorage.setItem('userData', JSON.stringify(this.user))
        localStorage.setItem('token', JSON.stringify(this.token))
        this.router.navigateByUrl('/')
      }
    });
  }

  logOut() {
    this.user.id = ''
    this.user.username = ''
    this.token = null
    localStorage.clear()
    this.router.navigateByUrl('/auth')
  }
}
