<%@ Page Language="C#" AutoEventWireup="true" CodeFile="Default.aspx.cs" Inherits="_Default" %>

<!DOCTYPE html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang=""> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8" lang=""> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9" lang=""> <![endif]-->
<!--[if gt IE 8]><!-->

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <title>Index</title>
    
<meta charset="utf-8">
<meta name="description" content="">
<meta name="viewport" content="width=device-width, initial-scale=1">

<link rel="stylesheet" href="css/bootstrap.min.css">
<link rel="stylesheet" href="css/flexslider.css">
<link rel="stylesheet" href="css/jquery.fancybox.css">
<link rel="stylesheet" href="css/main.css">
<link rel="stylesheet" href="css/responsive.css">
<link rel="stylesheet" href="css/animate.min.css">
<link rel="stylesheet" href="css/font-icon.css">
<link rel="stylesheet" href="https://cdn.staticfile.org/font-awesome/4.4.0/css/font-awesome.min.css">
</head>
<body>
    <form id="form1" runat="server">
   
<!-- header section -->
<section class="banner" role="banner">
<!--header navigation -->
  <header id="header">
    <div class="header-content clearfix"> <a class="logo" href="index.html"><img src="images/logo.png" alt=""></a>
      <nav class="navigation" role="navigation">
        <ul class="primary-nav">
          <li><a href="about.html">Login</a></li>
          <li><a href="contact.html">PAY</a></li>
		  <!--<li><a href="#">External Link</a></li>-->
        </ul>
      </nav>
      <a href="#" class="nav-toggle">Menu<span></span></a> </div>
      </header>
      <!--header navigation -->
  </section>
  <!-- header section -->
  <!-- text banner section -->
<section id="banner" class="banner no-padding">
  <div class="container-fluid">
    <div class="row no-gutter">
      <div class="flexslider">
        <ul class="slides">
          <li>
            <div class="col-md-12">
              <blockquote>
              <%--  <h1>Flowxa</h1>--%>
               <p id="txtgoldwords"></p>
              </blockquote>
            </div>
          </li>
          <li>
            <div class="col-md-12">
              <blockquote>
               <%-- <h1>Flowxa</h1>--%>
                <p id="txtgoldwords2"></p>
              </blockquote>
            </div>
          </li>
        </ul>
      </div>
     </div>
  </div>
</section>
<!-- Text banner section --> 
<!-- portfolio section -->
<section id="works" class="works section no-padding">
  <div class="container-fluid">
    <div class="row no-gutter">
      <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="http://go.flowx.cn" class="work-box"> <img src="images/work-1.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>说话的艺术</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>
      <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="http://cc.flowx.cn/BMI.aspx" class="work-box"> <img src="images/work-2.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>BMI</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>
    <%--  <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-3.jpg" class="work-box"> <img src="images/work-3.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
     <%-- <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-4.jpg" class="work-box"> <img src="images/work-4.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
     <%-- <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-5.jpg" class="work-box"> <img src="images/work-5.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
     <%-- <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-6.jpg" class="work-box"> <img src="images/work-6.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
     <%-- <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-7.jpg" class="work-box"> <img src="images/work-7.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
     <%-- <div class="col-lg-3 col-md-6 col-sm-6 work"> <a href="images/work-8.jpg" class="work-box"> <img src="images/work-8.jpg" alt="">
        <div class="overlay">
          <div class="overlay-caption">
            <h5>Project Name</h5>
            <p><i class="fa fa-search-plus fa-2x"></i></p>
          </div>
        </div>
        <!-- overlay --> 
        </a> </div>--%>
    </div>
  </div>
</section>
<!-- portfolio section --> 
  <!-- footer -->
<footer class="section footer">
  <div class="footer-bottom">
    <div class="container">
      <div class="col-md-12">
        <p>
        <ul class="footer-share" style="display:none;">
          <li><a href="#"><i class="fa fa-facebook"></i></a></li>
          <li><a href="#"><i class="fa fa-twitter"></i></a></li>
          <li><a href="#"><i class="fa fa-linkedin"></i></a></li>
          <li><a href="#"><i class="fa fa-google-plus"></i></a></li>
          <li><a href="#"><i class="fa fa-pinterest-p"></i></a></li>
          <li><a href="#"><i class="fa fa-vimeo"></i></a></li>
        </ul>
        </p>
        <p>©2024 Flowxa  浙ICP备17039044号-1<br>
        </p>
      </div>
    </div>
  </div>
</footer>
<!-- footer --> 
<!-- JS FILES --> 
<script src="https://cdn.staticfile.org/jquery/1.11.3/jquery.min.js"></script> 
<script src="js/bootstrap.min.js"></script> 
<script src="js/jquery.flexslider-min.js"></script> 
<script src="js/jquery.fancybox.pack.js"></script> 
<script src="js/jquery.waypoints.min.js"></script> 
<script src="js/retina.min.js"></script> 
<script src="js/modernizr.js"></script> 
<script src="js/main.js"></script>
    </form>

    <script>
		function Sub() {
			//document.forms['search'].submit();
            $("#form1").submit();
		}

        function getwords() {

            $.ajax({
                type: "post",
                url: "ajax/goldwords.ashx?act=2&idx=",
                async: false,
                cache: false,
                dataType: 'json',
                success: function (data) {
                  //  alert(data.formJson);
                    if (data.code == "ok") {
                        var elems = data.formJson;
                        $("#txtgoldwords").html(elems);
                        $("#txtgoldwords2").html(elems);
                        // window.location.href = elems;//"GoBackAPP";
                        //var t = setTimeout("goapp()", 5000)

                    } else {
                        //setalter(data.msg);
                    }
                }
            });

        }
        getwords();

    </script>
</body>
</html>
