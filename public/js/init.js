(function($){
  $(function(){

    $('.button-collapse').sideNav();

  }); // end of document ready
})(jQuery); // end of jQuery name space


var groups;
$.ajax({
            url: '/api/group',
            method: 'GET',
            success: function (data) {

                groups = data;
//                alert(JSON.stringify(data));
								for (var key in data.data)
                {
                  console.log(data.data[key]);
                }

              var GroupsVue = new Vue({
                el: '#app',
                data: {
                    response: groups
                },
                methods: {

                },
                ready: function () {
                  $('.modal-trigger').leanModal();
                }
              });

            },
            error: function (error) {
//                alert(JSON.stringify(error));
            }
});

