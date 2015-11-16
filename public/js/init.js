(function($){
  $(function(){

    $('.button-collapse').sideNav();
		console.log("ready");

  }); // end of document ready
})(jQuery); // end of jQuery name space


var vm = new Vue({
	el: '#app',
		data: {
			response: ''
		},
    ready: function () {

      // GET request
      this.$http.get('/api/group', function (data, status, request) {
          // set data on vm
          this.$set('response', data)
		      console.log("ready ready");
      }).error(function (data, status, request) {
          // handle error
					console.log(status)
      })
		$('.modal-trigger').leanModal();
		$('.tooltipped').tooltip({delay: 50});
		}
});


