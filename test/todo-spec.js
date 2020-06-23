

describe('angularjs homepage todo list', function() {
    it('should add a todo', function() {
      browser.get('https://angularjs.org');
	  browser.sleep(5000).then(
	  function(){
		 console.log("Waiting");
	  }
	)
      element(by.model('todoList.todoText')).sendKeys('write first protractor test');
	  browser.sleep(5000).then(
	   function(){
     console.log("Waiting");
	}
	)
      element(by.css('[value="add"]')).click();
	  browser.sleep(5000).then(
	  function(){
		 console.log("Waiting");
	  }
)
      var todoList = element.all(by.repeater('todo in todoList.todos'));
      expect(todoList.count()).toEqual(3);
      expect(todoList.get(2).getText()).toEqual('write first protractor test');
  
      // You wrote your first test, cross it off the list
      todoList.get(2).element(by.css('input')).click();
      var completedAmount = element.all(by.css('.done-true'));
      expect(completedAmount.count()).toEqual(2);
    });
  });
  
  