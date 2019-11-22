$ ->
  # new room
  Room.checkWs()
  window.room = new Room("ws://"+ window.location.host + window.location.pathname + "/chatting")

  room.ws_conn.onopen = room.joinRoom()

  room.ws_conn.onmessage = (e) ->
    room.reveiveMessage(e)

  # deal with message 
  $('#sayit-button').click ->
    text = $('#chat-form').val()
    if text
      sMessage = new Message("text", text)
      room.sendMessage(sMessage)
    else
      return
    
  $('#chat-form').keydown (e) ->
    if e.ctrlKey && e.keyCode == 13
      text = $('#chat-form').val()
      sMessage = new Message("text", text)
      room.sendMessage(sMessage)

class Room
  constructor: (ws_url) ->
    @ws_url = ws_url
    @ws_conn = new WebSocket(@ws_url)
    @userlist = []
    #@key = window.location.pathname.split('/')[2]

  @checkWs: ->
    unless window.WebSocket
      alert("you brower is not support websocket")
      return

  roomAddr: ->
    console.log @ws_url
  
  joinRoom: ->
    message = new Message("join", "#{@currentUser()} has join room")
    #console.log(JSON.stringify(message))
    # send to server join message

  currentUser: ->
    $("#user-name").text()

  sendMessage: (message) ->
    unless @ws_conn
      return
    @ws_conn.send(JSON.stringify(message))
    $('#chat-form').val('')

    # json websocket send
  reveiveMessage: (e) ->
      m= $.parseJSON(e.data)

      if m.Type == "join"
        @addUserToList(m.User)
      if m.Type == "leave"
        @removeUserFromList(m.User)

      console.log m
      rMessage = new TextMessage(m.Type, m.Text, m.User)
      rMessage.show()

  getUsersList: ->
      url = window.location.pathname + "/users.json"

      $.getJSON(url, (data) =>
          if data.Users != null
            @userlist = data.Users
      )
      return @userlist

  addUserToList: (user)->
    names = []
    $("#userlist span").each ->
      names.push $(this).text()

    console.log names
    if names.indexOf(user.Name) == -1
      $("#userlist>ul").append $("<li><img src=#{user.Avatar}/><span>#{user.Name}</span></li>")

  removeUserFromList: (user) ->
    $("#userlist span").each ->
      if $(this).text() == user.Name
        $(this).parent().remove()


# message user send to server
class Message
  constructor: (type, text) ->
    @type = type
    @text = text

  autoUrl: ->
    url = /(\b(https?|ftp|file):\/\/[-A-Z0-9+&@#\/%?=~_|!:,.;]*[-A-Z0-9+&@#\/%=~_|])/ig
    @text = @text.replace(url, "<a href=\"$1\" target='_blank'>$1</a>")
    return @
  
  textWrap: ->
    @text = @text.replace(/\n/g, "<br />")
    return @

# deal with text message, show with user
class TextMessage extends Message
  constructor: (@type, @text, @user) ->

  show: ->
    @autoUrl()
    @textWrap()
    $('.chat-main').append "<span class='message-avatar'><img src='#{@user.Avatar}'></span>[#{@user.Name}] " + @text + "<br>"

