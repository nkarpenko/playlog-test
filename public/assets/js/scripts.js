"use strict";

$(document).ready(function(){
  new PlaylogCodingTest();
});

class PlaylogCodingTest {
  constructor() {

    this._settings = {
      baseUrl: 'http://localhost:9000',

      // Selectors
      loginInfoSelector: '#login-info',
      loginFormSelector: '.login-form',
      loginButtonSelector: '#login-submit-btn',
      commentsSectionSelector: '#comments',
      commentFormWrapSelector: '#comment-form-wrap',
      commentSubmitSelector: '#comment-submit-form-selector',

    };

    this._buttonLiked = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-file-person-fill" viewBox="0 0 16 16"><path d="M12 0H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm-1 7a3 3 0 1 1-6 0 3 3 0 0 1 6 0zm-3 4c2.623 0 4.146.826 5 1.755V14a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-1.245C3.854 11.825 5.377 11 8 11z"/></svg>`;
    this._buttonUnliked = `<svg class="liked" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-hand-thumbs-up-fill" viewBox="0 0 16 16"><path d="M6.956 1.745C7.021.81 7.908.087 8.864.325l.261.066c.463.116.874.456 1.012.964.22.817.533 2.512.062 4.51a9.84 9.84 0 0 1 .443-.05c.713-.065 1.669-.072 2.516.21.518.173.994.68 1.2 1.273.184.532.16 1.162-.234 1.733.058.119.103.242.138.363.077.27.113.567.113.856 0 .289-.036.586-.113.856-.039.135-.09.273-.16.404.169.387.107.819-.003 1.148a3.162 3.162 0 0 1-.488.9c.054.153.076.313.076.465 0 .306-.089.626-.253.912C13.1 15.522 12.437 16 11.5 16H8c-.605 0-1.07-.081-1.466-.218a4.826 4.826 0 0 1-.97-.484l-.048-.03c-.504-.307-.999-.609-2.068-.722C2.682 14.464 2 13.846 2 13V9c0-.85.685-1.432 1.357-1.616.849-.231 1.574-.786 2.132-1.41.56-.626.914-1.279 1.039-1.638.199-.575.356-1.54.428-2.59z"/></svg>`;
  
    this.start();
  }

  start() {
    this._registerUiEvents();
  }

  _registerUiEvents(){

    // Init listeners.
    this._loginListener();
    this._commentListener();
  }

  _loginListener() {
    $(document).on('click', this._settings.loginButtonSelector, {Context: this}, this._login);
  }

  _login(event){
    let self = event.data.Context;

    event.preventDefault();

    console.log("do something _login");

    var name = $('#username').val();
    var url = self._settings.baseUrl + '/user/login/' + name;
 
    $.ajax({
      dataType: 'json',
      url: self._settings.baseUrl + '/user/login/' + name,
      type: "POST", 

      success:function(response){
        console.log(response);
      },
      complete:function(){
        alert('completed');

        return false;
      },
      error:function(req, err){ 
        // console.log('Error: ' + err);
      }
    });

    return false;
  }

  _commentListener(event) {
    $(document).on('click', this._settings.commentSubmitSelector, {Context: this}, this._createComment);
  }

  _createComment(event) {
    event.preventDefault();

    console.log("do something _comment");

  }

  _getAllComments() {

  }

  _likeComment() {

  }

  _loadComments() {

  }

}