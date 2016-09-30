var postHeaders = {
    'Content-Type': 'application/x-www-form-urlencoded'
};

function xfrm( a ) {
    var z = encodeURIComponent;
    var str = [];
    for ( var p in a )
        str.push( z( p ) + "=" + z( a[ p ] ) );
    return str.join( "&" );
};

angular.module( 'App', [ 'ngRoute' ] )
    .service("state", function(){
        var state;
        return {
          clear: function(){
            state = {};
          },
          get: function(){
            return state;
          }, set: function(a,s){
            state[a] = s;
          }
        }
    })
    .config( function ( $routeProvider ) {
        $routeProvider
            .when( "/", {
                templateUrl: "templates/login.html",
                controller: "LoginCtrl"
            } )
            .when("/register",{
              templateUrl:"templates/register.html",
              controller:"RegistrationCtrl"
            })
            .when( "/user/:id", {
                templateUrl: "templates/user.html",
                controller: "UserCtrl"
            } )
    } )
    .controller('RegistrationCtrl', function($scope, $http, $location){
      $scope.data = {
      };
      $scope.message = "";
      $scope.register = function(){
        $http({
          method:"POST",
          url:"http://localhost:8081/register",
          headers: postHeaders,
          transformRequest: xfrm,
          data: {
            "username": $scope.data.username,
            "elderly": $scope.data.elderly,
            "employee": $scope.data.employee,
            "student": $scope.data.student
          }
        }).then(
          function success(resp){
            $location.path("/").hash($scope.data.username);
          },
          function failure(resp){
              $scope.data.message = {
                content:resp.data.message,
                type: 'is-warning'
              };
          }
        )
      }
    })
    .controller( 'LoginCtrl', function ( $scope, $http, $location ) {
        $scope.data = {

        };

        $scope.login = function () {
            $http( {
                    method: "POST",
                    url: "http://localhost:8081/login",
                    headers: postHeaders,
                    transformRequest: xfrm,
                    data: {
                        "username": $scope.username
                    }
                } )
                .then( function success( resp ) {
                    $location.path( "/user/" + resp.data.id );
                }, function failure( resp ) {
                  $scope.data.message = {
                    content:resp.data.message,
                    type: 'is-warning'
                  };
                } );
        }
    } )
    .controller( 'UserCtrl', function ( $scope, $http, $routeParams, state ) {
      $scope.data = {

      };

      $scope.update = function(){
        $http({
          method:"GET",
          url:"http://localhost:8081/user/" + $routeParams.id
        }).then(function success(resp){
          $scope.data.user = resp.data.user;
          console.log($scope.data);
        }, function failure(resp){
          $scope.data.message = {
            content:resp.data.message,
            type: 'is-warning'
          };
        })
      };
      $scope.update();

        $scope.newPass = function () {
          if ($scope.data.order == undefined){
            $scope.data.order = {}
          }
        };

        $scope.CompleteOrder = function(){
          $http({
            method:"POST",
            url:"http://localhost:8081/user/" + $routeParams.id + "/pass/create",
            headers: postHeaders,
            transformRequest: xfrm,
            data: {
                "type": $scope.data.order.type,
                "payment": $scope.data.order.payment
            }
          }).then(function success(resp){
            console.log(resp);
            $scope.data.order = undefined;
            $scope.update();
          }, function failure(resp){
            $scope.data.message = {
              content:resp.data.message,
              type: 'is-warning'
            };
          })
        }
    } );
