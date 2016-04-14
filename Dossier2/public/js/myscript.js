
$(document).ready(function() {
	var camera = $('div#webcam #video'),
		requestApi = $('div#request-api'),
		labelsUl = $('ul#labels'),
		canvasWidth = camera.width(),
		canvasHeight = canvasWidth / 1.33;
		canvasDiv = '<canvas id="canvas" width="' + canvasWidth + '" height="' + canvasHeight + '"></canvas>';

	requestApi.append(canvasDiv);
	console.log("canvas size : " + canvasWidth + " + " + canvasHeight);

	var canvas = $('#canvas'),
		ctx = canvas.get(0).getContext('2d'),
		video = $('#video').get(0),
		videoObj = {
			'video' : true
		};

	function errBack(error) {
		console.log("Video capture error: ", error.code);
	};

	canvas.hide();

	if (!window.URL)
		window.URL = window.URL || window.webkitURL || window.msURL || window.oURL;
	if (!navigator.getUserMedia)
		navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;
	if (navigator.getUserMedia) {
		navigator.getUserMedia(videoObj, function(stream) {
			video.src = window.URL.createObjectURL(stream);
			video.play();
		}, errBack);
	}

	$('#snap').on("click", function() {
		ctx.drawImage(video, 0, 0, canvasWidth, canvasHeight);
		var imgBase64 = canvas.get(0).toDataURL('image/jpeg', 1);
		var dateNow = Date.now();
		//console.log(dataURL);
		labelsUl.empty();
		var myData = {
			'date' : Date(dateNow),
			'img' : imgBase64
		};
		// console.log(myData);

		$.post('/api/appengine', myData)
			.done(function() {
				$.getJSON('/api/appengine', function(data) {
					data[0].labelAnnotations.forEach(function(current) {
						labelsUl.append('<li class="result">' + current.description + ' - ' + Math.floor(current.score * 100) + '%</li>');
					}).done(function() {
						console.log( "JSON Get successed" );
					}).fail(function() {
						console.log("JSON Get failed");
					}).always(function() {
						console.log("get request sent");
					});
				});
			}).fail(function() {
				console.log("JSON Post failed");
			}).always(function() {
				console.log("post request sent");
			});
	});
});
