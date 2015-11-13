(function($){
  $(function(){

    $('.button-collapse').sideNav();

  }); // end of document ready
})(jQuery); // end of jQuery name space

// {"data":[{"name":"archie-sit","pools":[{"name":"/DMZ/audmzsblsitweb-80-pool","blue":["/DMZ/audmzsblsitweb-01:80","/DMZ/audmzsblsitweb-02:80"],"green":["/DMZ/audmzsblsitweb-03:80","/DMZ/audmzsblsitweb-04:80"],"state":"blue"}],"state":"blue"}]}

var groups;

$.ajax({
            url: '/api/group',
            method: 'GET',
            success: function (data) {
                groups = data;
                alert(JSON.stringify(data));
var GroupsVue = new Vue({
            el: '#app',
            data: {
                response: groups
            },
            methods: {

            },
            ready: function () {

            }
});
            },
            error: function (error) {
                alert(JSON.stringify(error));
            }
});

