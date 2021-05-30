// 调pride的API
$.ajaxSetup({
    headers:{
       'client': "clientforjson"
    }
 });

// 获取所有已经完成的任务，用来生成那个完成度任务的时间权
// function get_all_finished_tasks(){
//     $.get("/v1/pride", function(data, status){    
//         console.log("-------output pride----------")
//         console.log(data.memories)
//         console.log(status)
//         all_finished_tasks = data.memories
//         // data.memories的形式是[{Name:date,Alldays:{task1,task2,task3,……}]
//         console.log(all_finished_tasks[0].Name)
//         console.log(all_finished_tasks[0].Alldays)
//     })
// }

// get_all_finished_tasks()


// ***********************************************************************
// 下面是在页面上生成目标
// ***********************************************************************
all_goals = []
all_project = {}
function get_all_goals(){
        $.get("/v1/goaljson", function(data, status){
        // yourgoals = []
        // alert("Data: " + data + "\nStatus: " + status);
        console.log("输出data")
        console.log(data);
        console.log("输出data.goals")
        console.log(data.goals)
        var goal_collection = data.goals
        for (let i = 0; i < goal_collection.length; i++) {
            all_goals.push(goal_collection[i].name)
            
        }
        console.log(`i am printint the all_goals ${all_goals}`)
        // create_goalllist_div(data.goals);
        // use_chosen_style()
        })
      }

// get_all_goals()

// 生成每个目标的无序列表div
function make_ul(array) {
    // 创建ul标签
    var list = document.createElement('ul');
    console.log("i am printing the array")
    console.log(array[1])
    for (var i = 0; i < array.length; i++) {
        // 创建li标签
        // console.log("--------i am in the for functoin---------")
        list.setAttribute("id","all_goals_list")
        var item1 = document.createElement('li');
        // 创建div、span标签 
        var item2 = document.createElement("div");
        var item3 = document.createElement("span");
        // 生成list div
        item3.textContent = array[i];
        item2.appendChild(item3);
        item1.appendChild(item2);
        // Add it to the list:
        list.appendChild(item1);
    }
   
    // Finally, return the constructed list:
    // console.log(`我正在输出list${list}`)
    // console.log(list)
    return list;
}

$('span').bind('dblclick',
    function(){
        $(this).attr('contentEditable',true);
    });

// ***********************************************************************
// 下面是创建目标
// ***********************************************************************

// 下面的代码是点击添加目标的时候出现添加目标的选项

$(document).on("click","#click_add_input_goal_div",function(){
    var input_goal_area  =`<input id="add_goal_text" type="text" placeholder="Please input your goal">
    <select name="goal_socer_select" id="goal_socer_select">
     <option value="5">重要等级5</option>
     <option value="4">重要等级4</option>
     <option value="3">重要等级3</option>
     <option value="2">重要等级2</option>
     <option value="1">重要等级1</option>
     </select>
    <button id="add_goal" class="add_goal">添加目标</button>`
    var add_goal_area = document.getElementById("click_add_input_goal_div")
    add_goal_area.insertAdjacentElement("afterend",createNode(input_goal_area))
    add_goal_area.outerHTML = ""
})


// 下面的代码是点击创建目标这个button的时候，开始创建一个目标
var priority

$("#goal_socer_select").chosen().change(function(Event,eventobje){
    priority = eventobje.selected
})

$(document).on("click","#add_goal",function(){
    var goal = $("#add_goal_text").val()
    var your_created_goal = {'goal':goal,'priority':priority}
    $.ajax({
        type: "POST",
        url: "/v1/creategoal",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify(your_created_goal),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) {
            console.log("我已经成功创建了目标")
            get_all_unfinished_tasks_list()
        },
        failure: function (errMsg) {
          alert(errMsg);
        }
})
})

// create_goal()

// 更新目标

// var update_goal = {'goal':goal,'goalcode':id,'priority':priority,"goalstatus":goalstatus,"plantime":plantime,"finishtime":finishtime,"timerange":planmonth}
function create_goal(){
    $.ajax({
        type: "POST",
        url: "/v1/updategoal",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify(your_created_goal),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) {
            console.log("输出创建的goals")
            console.log(data)
        },
        failure: function (errMsg) {
          alert(errMsg);
        }
})
}


// ***********************************************************************
// 下面是创建goal下所有project and task的代码
// ***********************************************************************

var all_unfinished_tasks = []

var all_unfinished_tasks_tree = {}

function get_all_unfinished_tasks_list(){
    $.get("/v1/goalsgraph", function(data, status){
    // alert("Data: " + data + "\nStatus: " + status);
    console.log("~~~~~~~~~~~~~~~我正在打印所有未完成的任务~~~~~~~~~")
    console.log(data);
    // 提取所有的Goal = [{goal1的所有task},{goal2的所有task},{}]
    all_unfinished_tasks = data.goals;
    console.log("I AM PRINTING THE ALL UNFINISHED TASK")
    console.log(all_unfinished_tasks)
    if (all_unfinished_tasks == null) {
        return
    }
    else{
         // return all_unfinished_tasks;
    document.getElementById("goal_shown_and_edit").innerHTML = ""
    create_goal_project_task_div()
    }
   
   

// =======================================================
// 这是一个浮层的js部分，来自codepen
// =======================================================
jQuery(document).ready(function($){
	//open popup
	$('.cd-popup-trigger').on('click', function(event){
		event.preventDefault();
		$('.cd-popup').addClass('is-visible');
	});
	
	//close popup
	$('.cd-popup').on('click', function(event){
		if( $(event.target).is('.cd-popup-close') || $(event.target).is('.cd-popup') ) {
			event.preventDefault();
			$(this).removeClass('is-visible');
		}
	});
    //close popup when clicking the esc keyboard button
    document.onkeydown = function (evt) {
        if (evt.keyCode == 27) evt.preventDefault();
        $('.cd-popup').removeClass('is-visible');
    }

	// $(document).keyup(function(event){
    // 	if(event.which=='27'){
    // 		$('.cd-popup').removeClass('is-visible');
	//     }
    // });
});

    })
    
}

get_all_unfinished_tasks_list()

// 下面解决如何得到1个Goal下的所有Project

function unique(arr){
    return Array.from(new Set(arr))
}

function get_no_repeat_project(array){
    var project = []
    // 得到所有Project
    for (let i = 0; i < array.length; i++) {
        var tempproject = array[i].project
        project.push(tempproject)
    }
    // 对Project去重
    project = unique(project)
    return project
}

// 哈希处理的函数
String.prototype.hashCode = function(){
    if (Array.prototype.reduce){
        return this.split("").reduce(function(a,b){a=((a<<5)-a)+b.charCodeAt(0);return a&a},0);              
    } 
    var hash = 0;
    if (this.length === 0) return hash;
    for (var i = 0; i < this.length; i++) {
        var character  = this.charCodeAt(i);
        hash  = ((hash<<5)-hash)+character;
        hash = hash & hash; // Convert to 32bit integer
    }
    return hash;
}




// 生成goal-project-task三层的列表div
function create_goal_project_task_div(){
    // get_all_unfinished_tasks_list()
    // 创建all_goals的列表
    // console.log("---------print all_unfinished_tasks-----------")
    // console.log(all_unfinished_tasks)
    var goal_ul = document.createElement("ul")
    // project = [];
    // for (let i = 0; i < all_unfinished_tasks.length; i++) {
    //     var this_goal_name = all_unfinished_tasks.Name
    //     var projects_in_this_goal = all_unfinished_tasks[i].Allprojectsingoal
    //     for (let j = 0; j < projects_in_this_goal.length; j++) {
    //         project.push(projects_in_this_goal[j].Name)
    //     }
    //     console.log("+++++++++我正在输出Project++++++++++")
    //     console.log(project)
    // }


    //第一层循环，对所有的Goal 
    for (let i = 0; i < all_unfinished_tasks.length; i++) {
        // all_goals.push(all_unfinished_tasks[i].Name)
        // get one goal
        var temp_goal =  all_unfinished_tasks[i].Name
        var goal_priority = all_unfinished_tasks[i].Priority
        // console.log("!!!!!!!!! printint goal !!!!!!!!!!")
        // console.log(temp_goal)
        // var project = get_no_repeat_project(all_tasks_in_one_project)
        // 输出当下Goal里的所有Projects到一个array
        var project = []
        var projects_in_this_goal = all_unfinished_tasks[i].Allprojectsingoal
        for (let j = 0; j < projects_in_this_goal.length; j++) {
            project.push(projects_in_this_goal[j].Name)
        }
        // console.log("!!!!!!!!! printint project !!!!!!!!!!")
        // console.log(projects_in_this_goal)

        // 输出当下Goal里的所有task到一个

        // var all_tasks_in_one_project = all_unfinished_tasks[i].Alltasksingoal
        // console.log("!!!!!!!!! printint alltaskinonegoal !!!!!!!!!!")
        // console.log(all_tasks_in_one_project)
        var goalcode = all_unfinished_tasks[i].Goalcode.split(" ")[0]

        // 创建Goal的DIV
        var goal_li = document.createElement("li")
        goal_li.setAttribute("class","goal_li")
        var goal_span = document.createElement("span")
        goal_span.textContent = temp_goal
        goal_span.setAttribute("goal_code", `${goalcode}`)
        goal_span.setAttribute("id",`${goalcode}`)
        var add_project2goal_button = document.createElement("input")
        add_project2goal_button.setAttribute("class","add_project2goal_button")
        add_project2goal_button.setAttribute("id",`${temp_goal}`.hashCode())
        add_project2goal_button.setAttribute("type","image")
        add_project2goal_button.setAttribute("width","12")
        add_project2goal_button.setAttribute("height","12")
        add_project2goal_button.setAttribute("src","https://test-1255367272.cos.ap-chengdu.myqcloud.com/plus.svg")
        
        
        var finish_goal = document.createElement("button");
        finish_goal.setAttribute("class","finish_goal_class");
        finish_goal.setAttribute("goalcode",`${goalcode}`);
        finish_goal.setAttribute("goalname",`${temp_goal}`);
        finish_goal.setAttribute("priority",`${goal_priority}`);
        finish_goal.innerHTML = finish_goal_project_button;

        var giveup_goal = document.createElement("button");
        giveup_goal.setAttribute("class","giveup_goal_class");
        giveup_goal.setAttribute("goalcode",`${goalcode}`);
        giveup_goal.setAttribute("goalname",`${temp_goal}`);
        giveup_goal.setAttribute("priority",`${goal_priority}`);
        giveup_goal.innerHTML = giveup_goal_project_button;
        
        goal_li.appendChild(goal_span)
        goal_li.appendChild(add_project2goal_button)
        if (project.length == 0) {
            goal_li.appendChild(finish_goal)
            goal_li.appendChild(giveup_goal) 
        }
       
        // goal_li.appendChild(project_input)

        // 第二层循环，对第一个Goal下的所有Projects
        for (let j = 0; j < project.length; j++) {
            
             // 得到一个project下的所有tasks
            var all_tasks_in_one_project =  all_unfinished_tasks[i].Allprojectsingoal[j].Alltasksinproject
            // console.log("!!!!!!!!! printint all tasks in one project !!!!!!!!!!")
            // console.log(all_tasks_in_one_project)

            var project_ul = document.createElement("ul")
            var project_li = document.createElement("li")
            project_li.setAttribute("class","project_li")
            var project_span = document.createElement("span")
            var add_task2project_button = document.createElement("input")
            add_task2project_button.setAttribute("class","add_task2project_button")
            add_task2project_button.setAttribute("id",`${temp_goal}_${project[j]}`.hashCode())
            add_task2project_button.setAttribute("type","image")
            add_task2project_button.setAttribute("width","12")
            add_task2project_button.setAttribute("height","12")
            add_task2project_button.setAttribute("src","https://test-1255367272.cos.ap-chengdu.myqcloud.com/plus.svg")
            
            var finish_project = document.createElement("button");
            finish_project.setAttribute("class","finish_project_class");
            finish_project.setAttribute("goalcode",`${goalcode}`)
            finish_project.setAttribute("projectname",`${project[j]}`)
            finish_project.innerHTML = finish_goal_project_button;
    
            var giveup_project = document.createElement("button");
            giveup_project.setAttribute("class","giveup_project_class");
            giveup_project.setAttribute("goalcode",`${goalcode}`)
            giveup_project.setAttribute("projectname",`${project[j]}`)
            giveup_project.innerHTML = giveup_goal_project_button;
            
            project_span.textContent =project[j]
            project_li.appendChild(project_span)
            project_li.appendChild(add_task2project_button)

            if (all_tasks_in_one_project == null) {
                project_li.appendChild(finish_project)
                project_li.appendChild(giveup_project) 
            }

            project_ul.appendChild(project_li)
            goal_li.appendChild(project_ul)

            var task_ul = document.createElement("ul");
            task_ul.setAttribute("id",`${temp_goal}_${project[j]}`)
        //get one project

       
        // 第三层循环，对一个Project下的所有tasks
        if ( all_tasks_in_one_project == null ) {
            
        }
        else{

        
           for (let z = 0; z < all_tasks_in_one_project.length; z++) {
            //   if( all_tasks_in_one_project[z].project == project[j] ){
                var task_id = all_tasks_in_one_project[z].ID
                var parentid = all_tasks_in_one_project[z].parentid
                var tasktags = all_tasks_in_one_project[z].tasktags
                var obj = JSON.parse(tasktags)
                var urgenttag = obj["urgenttag"]
                var keyproblemtag = obj["keyproblemtag"]
                
                // var urgenttag = obj[""]
                var task_li = document.createElement("li")
                task_li.setAttribute("id",`${task_id}`)
                task_li.setAttribute("class","task_li")
                task_li.setAttribute("parentid",`${parentid}`)
                task_li.setAttribute("urgenttag",`${urgenttag}`)
                var task_div = document.createElement("div")
                var task_span = document.createElement("span")
                task_span.setAttribute("urgenttag",`${urgenttag}`)
                task_span.setAttribute("keyproblemtag",`${keyproblemtag}`)
                var add_sontask_button = document.createElement("input")
                add_sontask_button.setAttribute("class","add_sontask_button")
                add_sontask_button.type = "image"
                add_sontask_button.width = "12"
                add_sontask_button.height = "12"
                add_sontask_button.setAttribute("src","https://test-1255367272.cos.ap-chengdu.myqcloud.com/plus.svg")

                var add_task2today_button = document.createElement("button")
                add_task2today_button.setAttribute("class","add_task2today_button")
                // add_task2today_button.textContent = "Today"
                add_task2today_button.innerHTML = today_word;

                var add_task2tomorrow_button = document.createElement("button")
                add_task2tomorrow_button.setAttribute("class","add_task2tomorrow_button")
                add_task2tomorrow_button.innerHTML = tomorrow_word

                var giveup_task_button = document.createElement("button")
                giveup_task_button.setAttribute("class","giveup_task_button")
                // giveup_task_button.textContent = "Giveup"
                 
                giveup_task_button.innerHTML = giveup_word;
                task_span.textContent = all_tasks_in_one_project[z].task
                task_div.appendChild(task_span)
                // task_div.appendChild(add_sontask_button)
                task_div.appendChild(add_task2today_button)
                task_div.appendChild(add_task2tomorrow_button)
                task_div.appendChild(giveup_task_button)   
                task_li.appendChild(task_div)
                task_ul.appendChild(task_li)
            //   }
            }
        }
           project_li.appendChild(task_ul)
           goal_li.appendChild(project_ul)
        // temp_project[`${temp_project}`] = temp_tasks
        }
        // all_unfinished_tasks_tree[`${temp_goal}`] = temp_project
        
        goal_ul.appendChild(goal_li)
    }
    document.getElementById("goal_shown_and_edit").appendChild(goal_ul)
}


// 为子任务和孙子任务添加退格CSS
function add_class2sontask(){
    var all_tasks_div = document.getElementById("goal_shown_and_edit")
    if (condition) {
    }
}

// 在节点之后插入dom的函数
function insertAfter(newNode, referenceNode) {
    referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
}

//===============================================================================//
// 添加Project的一系列函数
// 点击Goal右边的加号按键，在Goal下面出现创建Project的input窗口
var add_project_plusbutton
$(document).on("click","input.add_project2goal_button",function(){
    add_project_plusbutton = $(this);
    console.log(add_project_plusbutton);
    var goal_code = $(this).prev().attr("goal_code")
    // 下面两行代码实现隐藏后面的对号和叉号
    if ( $(`button[goalcode=${goal_code}]`) == null){
    }
    else{
        var need_hidden_button = $(`button[goalcode=${goal_code}]`)
        need_hidden_button.hide()
    }

    var project_input = document.createElement("input")
    project_input.setAttribute("goal_code",`${goal_code}`)
    project_input.setAttribute("class","add_project_area")
    project_input.setAttribute("placeholder","please input your project")
    console.log("-------------这里是goal_code------------------------")
    console.log(goal_code)
    var goal_name_hash = $(this).prev().text().hashCode()
    console.log(goal_name_hash)
    $(`#${goal_name_hash}`).after(project_input)
    var add_project_position = document.getElementById(goal_name_hash)
    add_project_position.outerHTML =""
    // insertAfter(add_project_position, project_input)
    // add_project_position.appendChild(project_input)
})

// 输入完新的Project之后，点击回车，添加新Project的动作
$(document).on("keypress",".add_project_area",function(event)
    {
        
      if (event.key === "Enter") {
        // alert("我在这里这里这里这里")
        var goal_code = $(this).attr("goal_code")

        var input_project = $(this).val()
        var projectinfo = {"projectname":`${input_project}`,"goalcode":`${goal_code}`}
        console.log("((((((((((((((正在输出Projcet)))))))))))")
        console.log(projectinfo)
        $.ajax({
            type: "POST",
            url: "/v1/createproject",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(projectinfo),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) {
            //   alert(data)
            console.log("添加Project成功")
            // document.getElementsByClassName("add_projcet_area").outerHTML = ""
            // $(".add_project_area").remove()
                remove_add_project_area()
                show_new_added_project(goal_code,input_project)
            },
            failure: function (errMsg) {
              alert(errMsg);
            }
            //  alert(errMsg);
          })
        }
    })


    // 去掉添加Project的Input窗口，然后还原回之前的加号
    function remove_add_project_area(){
        document.getElementsByClassName("add_project_area").outerHTML = ""
        $(".add_project_area").after(add_project_plusbutton);
        $(".add_project_area").remove();
      }
    
    // 不刷新页面显示出新添加的Project
    function show_new_added_project(goal_code,added_project){

            var goal = document.getElementById(`${goal_code}`).textContent
            var project_ul = document.createElement("ul")
            var project_li = document.createElement("li")
            project_li.setAttribute("class","project_li")
            var project_span = document.createElement("span")
            var add_task2project_button = document.createElement("input")
            add_task2project_button.setAttribute("class","add_task2project_button")
            add_task2project_button.setAttribute("id",`${goal}_${added_project}`.hashCode())
            add_task2project_button.setAttribute("type","image")
            add_task2project_button.setAttribute("width","12")
            add_task2project_button.setAttribute("height","12")
            add_task2project_button.setAttribute("src","https://test-1255367272.cos.ap-chengdu.myqcloud.com/plus.svg")
            project_span.textContent = added_project
            project_li.appendChild(project_span)
            project_li.appendChild(add_task2project_button)
            project_ul.appendChild(project_li)
            var goalbutton = document.getElementById(`${goal_code}`).nextSibling
            goalbutton.insertAdjacentElement("afterend",project_ul)
            // var goalbutton = document.getElementById(`${goal_code}`).appendChild(project_ul)
            // insertAfter(goalbutton,project_ul)
            // alert('该函数是否发生调用')
    }

// 点击project右边的加号按键，在project后面添加输入任务的窗口
$(document).on("click","input.add_task2project_button",function(){
     
    console.log(`我正在输出onedit变量${onedit}`)
    console.log("正在执行点击加号的事件")
    if (onedit == true){
      alert("请先消除任务编辑框")
      return;
    }
    if (onedit == false){
      console.log("正在执行把onedit变成true的if语句")
      onedit = true //进入编辑模式
     //  console.log(`我正在if里输出onedit变量${onedit}`)
    }
     parentid = "unspecified"
     selected_project = $(this).prev().text()
     selected_goal = $(this).parent().parent().parent()["0"].firstChild.innerHTML
    //  如果后面有对号和叉号，不进行显示
     if ( $(`button[projectname=${selected_project}]`) == null){
        
    }
    else{
        var need_hidden_button = $(`button[projectname=${selected_project}]`)
        need_hidden_button.hide()
    }


     console.log("-------------------我正在输出project========")
     console.log(selected_project)
     console.log(selected_goal)
    //  console.log(selected_goal.firstChild.innerHTML)
    // console.log(createNode(add_task_div))
    // var add_task_div = createNode(add_task_div)
    var temp_selected_id = `${selected_goal}_${selected_project}`.hashCode()
    console.log(temp_selected_id)

    var add_plus = document.getElementById(temp_selected_id);

    var add_task_div_right_side = 
    `<li id="add_task_item_right_side">
    <div class="add_task_area" id="add_task" >
      <div class="input_task_area" id="input_task_area">
        <div class="task_container">
        <input class="input_text" type="text" id="noid" placeholder="please input your task" >
        </div>
        <div class="task_property_area">
          <div class="calendar_project_area">
            <div class="calendar_area" id="calendar_area">
              <button id="calendar" class="clock_button">
              ${calendar_button}
              </button>
            </div>
            <div id="goal_list" class="goal_list">
            </div>
            </div>
          <div id="tags_review_area" class="tags_review_area">
            <div id="tasktags_area">
              <button id="tags_button" class="tags_button">
                ${tags_button}
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="add_task_button_area" id="add_task_button_area")>
          <button class="add_task_click_button">添加任务</button>
          <button class="cancel_add_task_button">取消</button>
          <div class="add_tasktags_area" id="add_tasktags_area">
          </div>
      </div>
    </div>
  </li>`
    // $(temp_selected_id).append("<a>test the task append</a>");
    // var div = document.createElement("div");
    // 在获取到的div后面插入任务输入窗口，这个方法也可以用在子任务上面
    add_plus.insertAdjacentElement('afterend',createNode(add_task_div_right_side))
  })

//===============================================================================//

// 点击任务右边的加号按键，在任务下面添加输入子任务的窗口
$(document).on("click","input.add_sontask_button",function(){
     
    console.log(`我正在输出onedit变量${onedit}`)
    console.log("正在执行点击加号的事件")
    if (onedit == true){
      alert("请先消除任务编辑框")
      return;
    }
    if (onedit == false){
      console.log("正在执行把onedit变成true的if语句")
      onedit = true //进入编辑模式
     //  console.log(`我正在if里输出onedit变量${onedit}`)
    }
     parentid = $(this).parent().parent().attr("id")
     selected_project = $(this).parent().parent().parent().parent()["0"].firstChild.innerHTML
     selected_goal = $(this).parent().parent().parent().parent().parent().parent()["0"].firstChild.innerHTML
     
     console.log("-------------------我正在输出id project goal========")
    console.log(parentid)
     console.log(selected_project)
     console.log(selected_goal)

    var add_plus = document.getElementById(parentid);

    var add_task_div_right_side = 
    `
    <li id="add_task_item_right_side">
    <div class="add_task_area" id="add_task" >
      <div class="input_task_area" id="input_task_area">
        <div class="task_container">
        <input class="input_text" type="text" id="noid" placeholder="please input your task" >
        </div>
        <div class="task_property_area">
          <div class="calendar_project_area">
            <div class="calendar_area" id="calendar_area">
              <button id="calendar" class="clock_button">
              ${calendar_button}
              </button>
            </div>

            <div id="goal_list" class="goal_list">
            </div>
            </div>

          <div id="tags_review_area" class="tags_review_area">
            
            <div id="tasktags_area">
              <button id="tags_button" class="tags_button">
                ${tags_button}
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="add_task_button_area" id="add_task_button_area")>
          <button class="add_task_click_button">添加任务</button>
          <button class="cancel_add_task_button">取消</button>
          <div class="add_tasktags_area" id="add_tasktags_area">
          </div>
      </div>
    </div>
  </li>`
    // $(temp_selected_id).append("<a>test the task append</a>");
    var div = document.createElement("div");
    add_plus.insertAdjacentElement('afterend',createNode(add_task_div_right_side))
  })


// 尽量少地进行数据交互，把生成的子任务添加到tasktree里面
function show_new_added_task(inputedtask,task_id,parentid,goal,project){
    //  var task_id = all_tasks_in_one_project[z].ID
                // var parentid = all_tasks_in_one_project[z].parentid
                var task_li = document.createElement("li")
                task_li.setAttribute("id",`${task_id}`)
                task_li.setAttribute("class","task_li")
                task_li.setAttribute("parentid",`${parentid}`)
                var task_div = document.createElement("div")
                var task_span = document.createElement("span")

                var add_sontask_button = document.createElement("input")
                add_sontask_button.setAttribute("class","add_sontask_button")
                add_sontask_button.type = "image"
                add_sontask_button.width = "12"
                add_sontask_button.height = "12"
                add_sontask_button.setAttribute("src","https://test-1255367272.cos.ap-chengdu.myqcloud.com/plus.svg")

                var add_task2today_button = document.createElement("button")
                add_task2today_button.setAttribute("class","add_task2today_button")
                // add_task2today_button.textContent = "Today"
                add_task2today_button.innerHTML = today_word;
               
                var add_task2tomorrow_button = document.createElement("button")
                add_task2tomorrow_button.setAttribute("class","add_task2tomorrow_button")
                add_task2tomorrow_button.innerHTML = tomorrow_word;
               
                var giveup_task_button = document.createElement("button")
                giveup_task_button.setAttribute("class","giveup_task_button")
                // giveup_task_button.textContent = "Giveup"
                giveup_task_button.innerHTML = giveup_word;
                
                task_span.textContent = inputedtask
                task_div.appendChild(task_span)
                // task_div.appendChild(add_sontask_button)
                task_div.appendChild(add_task2today_button)
                task_div.appendChild(add_task2tomorrow_button)
                task_div.appendChild(giveup_task_button)   
                task_li.appendChild(task_div)
                document.getElementById(`${goal}_${project}`).appendChild(task_li)

}


// 右边点击明天按键把任务调度到明天
$(document).on("click",".add_task2tomorrow_button",function(){
    API_scr = "/v1/tomorrowjson"
    var clicked_taskid = []
    clicked_taskid.push($(this).parent().parent().attr("id"))
    var task_li_div = document.getElementById(clicked_taskid)
    // var giveup_taskids = []
    // giveup_taskids.push(clicked_taskid)
    // var sons = sontree[clicked_taskid]
    // for (let i = 0; i < sons.length; i++) {
    //   var grandson_id = sons[i].ID
    //   giveup_taskids.push(grandson_id)
    //   grandsons = sontree[grandson_id]
    //   for (let j = 0; j < grandsons.length; j++) {
       
    //     giveup_taskids.push(grandsons[j].ID)
    //   }
    // }
    console.log("==========这里是所有要添加到明天任务的ID=============")
    console.log(clicked_taskid)
    // console.log(giveup_taskids)
    var info = {
        'tomorrowtaskids':clicked_taskid,
    }
          // console.log(updatedinfo)
          $.ajax({
            type: "POST",
            url: "/v1/tomorrowtasksbatch",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(info),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) { 
              console.log("=======add task2tomorrow successful=========")
              console.log(data)
              // alert(data); 
            //   geteverydaytask()
            // alert("您已经成功把任务添加到明天")
            task_li_div.outerHTML = ""
            show_tomorrow_tree()
            },
            failure: function (errMsg) {
              console.log("this is erro")
              alert(errMsg);
  
            }
          })   
  })

// 左边todaytree点击明天的按键把任务添加到明天  todaytree_add_task2tomorrow_button
$(document).on("click",".left_add_task2tomorrow_button",function(){
    // 先判断是今天还是明天，然后决定是把任务更新到今天，还是更新到明天
    
        var clicked_taskid = []
        clicked_taskid.push($(this).attr("value"))
        var task_li_div = document.getElementById(clicked_taskid)
        console.log("==========这里是所有要添加到明天任务的ID=============")
        console.log(clicked_taskid)
        // console.log(giveup_taskids)
        var info = {
            'tomorrowtaskids':clicked_taskid,
        }
              // console.log(updatedinfo)
              $.ajax({
                type: "POST",
                url: "/v1/tomorrowtasksbatch",
                // The key needs to match your method's input parameter (case-sensitive).
                data: JSON.stringify(info),
                contentType: "application/json; charset=utf-8",
                dataType: "json",
                success: function (data) { 
                  console.log("=======add task2tomorrow successful=========")
                  console.log(data)
                  // alert(data); 
                //   geteverydaytask()
                // alert("您已经成功把任务添加到明天")
                task_li_div.outerHTML = ""
                // geteverydaytask()
                // show_tomorrow_tree()
                show_today_tree()
                },
                failure: function (errMsg) {
                  console.log("this is erro")
                  alert(errMsg);
      
                }
              })   
   
    
});




    $(document).on("click",".left_add_task2today_button",function(){
    
   
        // alert("我正在调用这个今天的函数")
        var clicked_taskid = $(this).attr("value")
        // alert(clicked_taskid)
        var task_li_div = document.getElementById(clicked_taskid)
        console.log("==========这里是所有要调度到今天任务的ID=============")
        console.log(clicked_taskid)
        // console.log(giveup_taskids)
        
        var updatedinfo =
        {'ifdissect': 'no', 
        'deadline': 'unspecified', 
        'starttime': 'unspecified', 
        'endtime': 'unspecified', 
        'tasktagsorigin': 'unspecified', 
        'parentid': 'unspecified', 
        'timedevotedto_a_task': 0, 
        'goalcode': 'xxx', 
        'client': 'gtdcli', 
        'taglight': 'no', 
        'note': 'unspecified', 
        'reviewalgolight': 'no', 
        'reviewalgo': {}, 
        'parentproject': 'unspecified', 
        'id': clicked_taskid, 
        'place': 'unspecified', 
        'finishtime': 'unspecified', 
        'inbox': 'nocontent', 
        'project': 'inbox', 
        'plantime': 'today', 
        'taskstatus': 'unfinished', 
        'tasktags': {}
    }
            
              // console.log(updatedinfo)
              $.ajax({
                type: "POST",
                url: "/v1/update",
                // The key needs to match your method's input parameter (case-sensitive).
                data: JSON.stringify(updatedinfo),
                contentType: "application/json; charset=utf-8",
                dataType: "json",
                success: function (data) { 
                //   alert("我正在把任务调度到660606")
                //   alert(data)
                  
                  console.log(data)
                  // alert(data); 
                  task_li_div.outerHTML = ""
              show_today_tree()
            
                //   get_all_unfinished_tasks_list()
                  
                },
                failure: function (errMsg) {
                  console.log("this is erro")
                  alert(errMsg);
                }
              })
    

  })



// 点击左侧Todaytree的向右的键头，把任务日期调到660606这一天，同时移除左侧的任务
$(document).on("click",".todaytree_move_task2someday_button",function(){
    var clicked_taskid = $(this).attr("value")
    // alert(clicked_taskid)
    var task_li_div = document.getElementById(clicked_taskid)
    console.log("==========这里是所有要调度到今天任务的ID=============")
    console.log(clicked_taskid)
    // console.log(giveup_taskids)
    
    var updatedinfo =
    {'ifdissect': 'no', 'deadline': 'unspecified', 'starttime': 'unspecified', 'endtime': 'unspecified', 'tasktagsorigin': 'unspecified', 'parentid': 'unspecified', 'timedevotedto_a_task': 0, 'goalcode': 'xxx', 'client': 'gtdcli', 'taglight': 'no', 'note': 'unspecified', 'reviewalgolight': 'no', 'reviewalgo': {}, 'parentproject': 'unspecified', 'id': clicked_taskid, 'place': 'unspecified', 'finishtime': 'unspecified', 'inbox': 'nocontent', 'project': 'inbox', 'plantime': '200111', 'taskstatus': 'unfinished', 'tasktags': {}}
        
          // console.log(updatedinfo)
          $.ajax({
            type: "POST",
            url: "/v1/update",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(updatedinfo),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) { 
            //   alert("我正在把任务调度到660606")
            //   alert(data)
              
              console.log(data)
              // alert(data); 
              task_li_div.outerHTML = ""
            //   geteverydaytask()
            show_today_or_tomorrow_task()
              get_all_unfinished_tasks_list()
              
            },
            failure: function (errMsg) {
              console.log("this is erro")
              alert(errMsg);
            }
          })
          
})





// 点击放弃按键放弃任务
$(document).on("click",".giveup_task_button",function(){
    var clicked_taskid = []
    clicked_taskid.push($(this).parent().parent().attr("id"))
    var task_li_div = document.getElementById(clicked_taskid)
    // var giveup_taskids = []
    // giveup_taskids.push(clicked_taskid)
    // var sons = sontree[clicked_taskid]
    // for (let i = 0; i < sons.length; i++) {
    //   var grandson_id = sons[i].ID
    //   giveup_taskids.push(grandson_id)
    //   grandsons = sontree[grandson_id]
    //   for (let j = 0; j < grandsons.length; j++) {
       
    //     giveup_taskids.push(grandsons[j].ID)
    //   }
    // }
    console.log("==========这里是所有要放弃的ID=============")
    console.log(clicked_taskid)
    // console.log(giveup_taskids)
    var info = {
      "giveuptaskids":clicked_taskid,
    }
          // console.log(updatedinfo)
          $.ajax({
            type: "POST",
            url: "/v1/giveuptasksbatch",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(info),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) { 
              console.log("=======give up task successful=========")
              console.log(data)
              // alert(data); 
            //   geteverydaytask()
            task_li_div.outerHTML = ""
            },
            failure: function (errMsg) {
              console.log("this is erro")
              alert(errMsg);
  
            }
          })
          
  })

// 把任务调度到今天
$(document).on("click",".add_task2today_button",function(){
    // alert("调试到今天的事件被触发")
    var clicked_taskid = $(this).parent().parent().attr("id")
    var task_li_div = document.getElementById(clicked_taskid)

    console.log("==========这里是所有要调度到今天任务的ID=============")
    console.log(clicked_taskid)
    // console.log(giveup_taskids)
    
    var updatedinfo =
        {
            "ifdissect":"no",
            "deadline":"unspecified",
            "starttime":"unspecified",
            "endtime":"unspecified",
            "tasktagsorigin":"unspecified",
            "parentid":"unspecified",
            "timedevotedto_a_task":0,
            "goalcode":"xxx",
            "client":"gtdcli",
            "taglight":"no",
            "note":"unspecified",
            "reviewalgolight":"no",
            "reviewalgo":{},
            "parentproject":"unspecified",
            "id":clicked_taskid,
            "place":"unspecified",
            "finishtime":"unspecified",
            "inbox":"nocontent",
            "project":"inbox",
            "plantime":"today",
            "taskstatus":"unfinished",
            "tasktags":{}
             
        }
          // console.log(updatedinfo)
          $.ajax({
            type: "POST",
            url: "/v1/update",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(updatedinfo),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) { 
            //   alert("我成功把任务调度到今天")
              console.log(data)
              // alert(data); 
              task_li_div.outerHTML = ""
              show_today_tree()
            },
            failure: function (errMsg) {
              console.log("this is erro")
              alert(errMsg);
            }
          })
          
  })




// 生成存放goals的div的函数
function create_goals_div_to_html(){
    console.log("====================")
    goal_list = make_ul(all_goals)
    document.getElementById("goal_shown_and_edit").appendChild(goal_list)
}
// 点击页面上的“所有目标”把生成的goals_div放到html里面
$(document).on("click","#goal_manager",function(){
    // create_goals_div_to_html()
    create_goal_project_task_div()
})



// 点击目标完成的按键，完成目标
$(document).on("click",".finish_goal_class",function(){
        var goal_name = $(this).attr("goalname")
        var goal_status = "finished"
        var goal_code = $(this).attr("goalcode")
        var finish_time = today
        var plan_time = 201020
        var priority = $(this).attr("priority")
        var this_goal_div = $(this).parent()[0]
    
        var updatedinfo = {
            "goalcode":goal_code,
            "goalname":goal_name,
            "goalstatus":goal_status,
            "finishtime":finish_time,
            "plantime":plan_time,
            "priority":priority
        }
        console.log(updatedinfo)
        $.ajax({
            type: "POST",
            url: "/v1/updategoal",
            // The key needs to match your method's input parameter (case-sensitive).
            data: JSON.stringify(updatedinfo),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (data) { 
            //   alert("我正在把任务调度到今天")
              console.log(data)
              // alert(data); 
              this_goal_div.outerHTML = ""
            },
            failure: function (errMsg) {
              console.log("this is erro")
              alert(errMsg);
            }
          })
    })



// 点击目标放弃的按键，放弃目标
$(document).on("click",".giveup_goal_class",function(){
    alert("放弃目标需谨慎，你确定要放弃这个目标吗？")
    var goal_name = $(this).attr("goalname")
    var goal_status = "giveup"
    var goal_code = $(this).attr("goalcode")
    var finish_time = today
    var plan_time = 201020
    var priority = $(this).attr("priority")
    var this_goal_div = $(this).parent()[0]

    var updatedinfo = {
        "goalcode":goal_code,
        "goalname":goal_name,
        "goalstatus":goal_status,
        "finishtime":finish_time,
        "plantime":plan_time,
        "priority":priority
    }
    console.log(updatedinfo)
    $.ajax({
        type: "POST",
        url: "/v1/updategoal",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify(updatedinfo),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) { 
        //   alert("我正在把任务调度到今天")
          console.log(data)
          // alert(data); 
          this_goal_div.outerHTML = ""
        },
        failure: function (errMsg) {
          console.log("this is erro")
          alert(errMsg);
        }
      })
})

// 点击项目完成的按键，完成项目
$(document).on("click",".finish_project_class",function(){
    var project_name = $(this).attr("projectname")
    var goal_code = $(this).attr("goalcode")
    var projects_status = "finished"
    var this_project_div = $(this).parent().parent()[0]
    console.log(this_project_div)
    var updatedinfo = {
        "projectname":project_name,
        "goalcode":goal_code,
        "projectstatus":projects_status
    }
    console.log(updatedinfo)
    $.ajax({
        type: "POST",
        url: "/v1/updateproject",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify(updatedinfo),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) { 
        //   alert("我正在把任务调度到今天")
          console.log(data)
          // alert(data); 
          this_project_div.outerHTML = ""
        },
        failure: function (errMsg) {
          console.log("this is erro")
          alert(errMsg);
        }
      })
      

})



// 点击项目放弃的按键，放弃项目
$(document).on("click",".giveup_project_class",function(){
    var project_name = $(this).attr("projectname")
    var goal_code = $(this).attr("goalcode")
    var projects_status = "giveup"
    var this_project_div = $(this).parent().parent()[0]
    console.log(this_project_div)

    var updatedinfo = {
        "projectname":project_name,
        "goalcode":goal_code,
        "projectstatus":projects_status
    }
    console.log(updatedinfo)
    $.ajax({
        type: "POST",
        url: "/v1/updateproject",
        // The key needs to match your method's input parameter (case-sensitive).
        data: JSON.stringify(updatedinfo),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data) { 
        //   alert("我正在把任务调度到今天")
          console.log(data)
          // alert(data); 
          this_project_div.outerHTML = ""
        },
        failure: function (errMsg) {
          console.log("this is erro")
          alert(errMsg);
        }
      })
})



function getCookieValue(a) {
    var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)');
    return b ? b.pop() : '';
}








// 所有任务的echarts
function review_echarts(){
     // 基于准备好的dom，初始化echarts实例 http://echarts.baidu.com/examples/#chart-type-treemap
        var myChart = echarts.init(document.getElementById(' '),{width:"300px",height:"60px"});

        // 指定图表的配置项和数据
/*        var option = {
            title: {
                text: 'ECharts 入门示例'
            },
            tooltip: {},
            legend: {
                data:['销量']
            },
            xAxis: {
                data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
            },
            yAxis: {},
            series: [{
                name: '销量',
                type: 'bar',
 data: [5, 20, 36, 10, 10, 20]
}]
};

// 使用刚指定的配置项和数据显示图表。
myChart.setOption(option);
*/
//myChart.showLoading();
//$.get('data/asset/data/flare.json', function (data) {
//myChart.hideLoading();

jQuery.ajaxSetup({async:false});
  function  getprojectrees(){
    projectsfortree = []
    $.get("/v1/projectsjson", function(data, status){
    var datafromserver = data
    //console.log(indicators.hist)
    projects = datafromserver.projects
    // histv = indicators.hist
    // signalv = indicators.signal
    // kdjkv =  indicators.kdjk
    // kdjdv =  indicators.kdjd
    // kdjjv =  indicatof       rs.kdjj

    for(j=0;j<projects.length;j++){
        console.log("--------")
    //macdvalue.push(projects[i])
    singleproject_tasks = []
    for(i=0;i<projects[j].Alltasksinproject.length;i++){
     var alltasks = projects[j].Alltasksinproject
      singleproject_tasks.push({"name":alltasks[i].task +"-"+alltasks[i].ID.toString(),"value":223})
    }
    projectsfortree.push({"name":projects[j].Name,"children":singleproject_tasks})
    }
   })
return projectsfortree
}
var data = {"name":"dm","children":getprojectrees()}

    myChart.showLoading();
    myChart.hideLoading();
    var _zr = myChart.getZr();
    //  ShowObjProperty(_zr);
    _zr.add(new echarts.graphic.Text({
     style: {            
   x: _zr.getWidth() / 2,
   y: _zr.getHeight() / 2,
   color: '#666',
   text: '集团重大风险',
   textAlign: 'center',
   textFont : 'bold 20px verdana'
   }}  
    ));
/*
    echarts.util.each(data.children, function (datum, index) {
        index % 2 === 0 && (datum.collapsed = true);
    });

    myChart.setOption(option = {
        tooltip: {
            trigger: 'item',
            triggerOn: 'mousemove'
        },
        series: [
            {
                type: 'tree',

                data: [data],

                top: '2%',
                left: '7%',
                bottom: '1%',
                right: '20%',

                symbolSize: 7,

                label: {
                    normal: {
                        position: 'left',
                        verticalAlign: 'middle',
                        align: 'right',
                        fontSize: 15
                    }
                },

                leaves: {
                    label: {
                        normal: {
                            position: 'right',
                            verticalAlign: 'middle',
                            align: 'left'
                        }
                    }
                },

                expandAndCollapse: true,
                animationDuration: 550,
                animationDurationUpdate: 750
            }
        ]
    });
*/
myChart.setOption(option = {
    tooltip: {
        trigger: 'item',
        triggerOn: 'mousemove'
    },
    series: [
        {
            type: 'tree',

            data: [data],

            top: '18%',
            bottom: '14%',

            layout: 'radial',

            symbol: 'pin',

            symbolSize: 10,

            initialTreeDepth: 2,
            rom:true,
            emphasis:{
                lineStyle:{width:20}
            },
            label:{show:true,position:"right",rotate:12,fontStyle:'italic',fontSize:15},
            lineStyle:{width:1,color:"red",curveness:0.1},
            animationDurationUpdate: 750

        }
    ]
});
}
  

// review_echarts()


// var my2Chart = echarts.init(document.getElementById('reviewalgo_temp'));

    // 指定图表的配置项和数据
    // var option = {
    //     title: {
    //         text: 'ECharts 入门示例'
    //     },
    //     tooltip: {},
    //     legend: {
    //         data:['销量']
    //     },
    //     xAxis: {
    //         data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
    //     },
    //     yAxis: {},
    //     series: [{
    //         name: '销量',
    //         type: 'bar',
    //         data: [5, 20, 36, 10, 10, 20]
    //     }]
        
    // };

    // 使用刚指定的配置项和数据显示图表。
    // my2Chart.setOption(option);







// ---------------------------------------------------
// ----------------------------------------------------
var token = getCookieValue("email")
var  WebSocketurl = 'ws://47.100.100.141:777'
var localhostWebSocketurl = "ws://localhost:777"
var httpsWebSocketurl ='wss://www.blackboxo.top/wss'
var heartCheck = {
    timeout: 6000,//60ms
    timeoutObj: null,
    ws:null,
    serverTimeoutObj: null,
    reset: function(){
        clearTimeout(this.timeoutObj);
        clearTimeout(this.serverTimeoutObj);
　　　　 this.start();
    },
    start: function(){
        //alert(this.ws)
        var self = this;
        this.timeoutObj = setTimeout(function(){
        console.log("------我在这里测试心跳-----------")
        // alert(self.ws)
        console.log("------我在这里测试心跳-----------")
        self.ws.send("HeartBeat");
        self.serverTimeoutObj = setTimeout(function(){
        self.ws.close();//如果onclose会执行reconnect，我们执行ws.close()就行了.如果直接执行reconnect 会触发onclose导致重连两次
            }, self.timeout)
        }, this.timeout)
    },
}





function connect(url) {
  var ws = new WebSocket(url);
  
 //   ws.binaryType = 'arraybuffer';
    ws.onopen = function() {
    //alert(ws)
    heartCheck.ws = ws;
    // alert(heartCheck.ws)
    heartCheck.start()
    // subscribe to some channels
    console.log("发送数据到服务器进行测试")
    ws.send("Message to send");
    console.log("发送数据到服务器进行测试")
    // ws.send(JSON.stringify({
    //   "a":1
    //     //.... some message the I must send when I connect ....
    // }));
  };

  ws.onmessage = function(e) {
    heartCheck.reset();
    console.log('Message:', e.data);
    if (e.data =="服务器发送图像数据到前端"){
     //var data = {"name":"goals","children":getgoaltrees()}
    //  doooo(data)
     loadlink()
    }
   
    // var data = new Uint8Array(e.data);
    // player.feed(data);
    
  };
  
  //超市重新链接
  ws.onclose = function(e) {
    console.log('Socket is closed. Reconnect will be attempted in 1 second.', e.reason);
    setTimeout(function() {
    // connect('ws://localhost:777');
    connect(url);
    // connect('ws://localhost:777');
    }, 1000);
  };

  ws.onerror = function(err) {
    console.error('Socket encountered error: ', err.message, 'Closing socket');
    ws.close();
  };
}


//connect(localhostWebSocketurl)
 //connect('ws://localhost:777')
// connect(httpsWebSocketurl)
// connect(localhostWebSocketurl+"/email="+token)





function calculateMA(dayCount, data) {
    console.log("i had no ho")
    var result = [];
    for (var i = 0, len = data.length; i < len; i++) {
        if (i < dayCount) {
            result.push('-');
            continue;
        }
        var sum = 0;
        for (var j = 0; j < dayCount; j++) {
            sum += data[i - j];
        }
        result.push(+(sum / dayCount).toFixed(3));
    }
    console.log(result)
    return result;
    
}



//echarts.init(document.getElementById('main')).dispose();
var myChart = echarts.init(document.getElementById('reviewalgo_tem'));


function gettoday(){
    var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = yyyy.toString().substring(2,4)+mm+dd  ;
return today
}

//https://stackoverflow.com/questions/25446628/ajax-jquery-refresh-div-every-5-seconds

var resize = function() {
      myChart.resize({
        // width: window.innerWidth ,
        // height: window.innerHeight
        width: $(window).width()/2,
        height: "300px"
      });
    };

    resize();

    //窗口变动时自适应
    window.onresize = function() {
      resize();
    };

























     // 基于准备好的dom，初始化echarts实例 http://echarts.baidu.com/examples/#chart-type-treemap
      //  var myChart = echarts.init(document.getElementById('main'));
//https://stackoverflow.com/questions/133310/how-can-i-get-jquery-to-perform-a-synchronous-rather-than-asynchronous-ajax-re
//https://stackoverflow.com/questions/25488915/async-false-for-shorthand-ajax-get-request
jQuery.ajaxSetup({async:false});




function loadlink(){
var standard_group = ['totalscore',"todayscore",'doanimportantthingearly','markataskimmediately','alwaysprofit','patience','usebrain','battlewithlowerbrain','learnnewthings','makeofthingsuhavelearned','difficultthings','threeminutes','getlesson','learntechuse','thenumberoftasks_score','serviceforgoal_score'];
var datecategory = [];
var totalscore = [];
var todayscore = [];
var patience = [];
var alwaysprofit = [];
var usebrain = [];
var battlewithlowerbrain = [];
var learnnewthings = [];
var makeuseofthingsuhavelearned = [];
var difficultthings = [];
var threeminutes = [];
var getlesson = [];
var learntechuse = [];
var thenumberoftasks_score = [];
var serviceforgoal_score = [];
var onlystartatask_score = [];
var atomadifficulttask_score = []; 
var doanimportantthingearly = [];
var markataskimmediately = [];
var buildframeandprinciple = [];
var acceptfact = [];
var executeablity = [];


$.get("/v1/reviewdaydatajson", function(data, status){
       // alert("Data: " + data.reviewdata[0].ID + "\nStatus: " + status);
        for(i=0;i<data.reviewdata.length;i++){
            var dateofreview =  data.reviewdata[i].date

            if (dateofreview !="unspecified"){

            
            datecategory.push(data.reviewdata[i].date);
            
            if (data.reviewdata[i].details != ""){
               var obj_detailofreview = JSON.parse(data.reviewdata[i].details);
          totalscore.push(obj_detailofreview.totalscore);
          patience.push(obj_detailofreview["patience"]);
         // alert(obj_detailofreview.patience);
          usebrain.push(obj_detailofreview["usebrain"]);
         // alert(obj_detailofreview.usebrain);
          battlewithlowerbrain.push(obj_detailofreview.battlewithlowerbrain);
          learnnewthings.push(obj_detailofreview.learnnewthings);
          makeuseofthingsuhavelearned.push(obj_detailofreview.makeuseofthingsuhavelearned);
          difficultthings.push(obj_detailofreview.difficultthings);
          threeminutes.push(obj_detailofreview.threeminutes);
          executeablity.push(obj_detailofreview.executeability_score*obj_detailofreview.totalscore);
           getlesson.push(obj_detailofreview.getlesson);
           learntechuse.push(obj_detailofreview.learntechuse);
          alwaysprofit.push(obj_detailofreview.alwaysprofit);
          thenumberoftasks_score.push(obj_detailofreview.thenumberoftasks_score);
          serviceforgoal_score.push(obj_detailofreview.serviceforgoal_score);
           onlystartatask_score.push(obj_detailofreview.onlystartatask_score);
           atomadifficulttask_score.push(obj_detailofreview.atomadifficulttask);
        doanimportantthingearly.push(obj_detailofreview.doanimportantthingearly);
        markataskimmediately.push(obj_detailofreview.markataskimmediately);
        buildframeandprinciple.push(obj_detailofreview.buildframeandprinciple);
        acceptfact.push(obj_detailofreview.acceptfact)
            }else{

               totalscore.push(0);
          patience.push(0);
         // alert(obj_detailofreview.patience);
          usebrain.push(0);
         // alert(obj_detailofreview.usebrain);
          battlewithlowerbrain.push(0);
          learnnewthings.push(0);
          makeuseofthingsuhavelearned.push(0);
          difficultthings.push(0);
          threeminutes.push(0);
          executeablity.push(0);
           getlesson.push(0);
           learntechuse.push(0);
          alwaysprofit.push(0);
          thenumberoftasks_score.push(0);
          serviceforgoal_score.push(0);
           onlystartatask_score.push(0);
           atomadifficulttask_score.push(0);
        doanimportantthingearly.push(0);
        markataskimmediately.push(0);
        buildframeandprinciple.push(0);
        acceptfact.push(0)


            }
        }
         if (data.reviewdata[i].date===gettoday()){break;}
         
};    








//添加一条今日线
todayscore = new Array(totalscore.length).fill(totalscore[totalscore.length-1]);
// alert(todayscore)
//alert(serviceforgoal_score);
/*

alert(totalscore);
alert(datecategory);
alert(patience);
alert(makeuseofthingsuhavelearned);
alert(learnnewthings);
alert(usebrain);
alert(battlewithlowerbrain);

*/
}
);


//finance api
var thisyear = 10000
var left = 100090
$.get("/finance/getrewardleft", function(data, status){
    /*
    
    {"left":finance["left"],"available":40,"budget":90,"thisyear":finance["thisyear"],"code":200}
    */
    var k = JSON.parse(data)
        thisyear = k.thisyear
        left = k.left
        // alert(data)
        // alert(left)     
});  

// alert(thisyear)

myChart.showLoading();
myChart.hideLoading();
var datafromserver = ['yangming','is','181923','201942','周五','周六','周日'];
myChart.clear();
myChart.setOption(



option = {
    graphic: { // 一个图形元素，类型是 text，指定了 id。
            type: 'text',
            id: 'text1'
        },
    title: {
        text: 'Evaluation system'
    },


    graphic:{ // 一个图形元素，类型是 text，指定了 id。
            type: 'text',
            id: 'text1',
            style: {
            text: totalscore[totalscore.length - 1].toFixed(2),
            x: 10,
            y: 20
    }
        },


    tooltip : {
        trigger: 'axis',
        axisPointer: {
            type: 'cross',
            label: {
                backgroundColor: '#6a7985'
            }
        }
    },
  
    legend: {
        data:['Totalscore','MA5','MA10','Make use of things u have learned','The total score of  number of tasks finished','Start an important thing early']
    },
    toolbox: {
        feature: {
            saveAsImage: {}
        }
    },
    // grid: {
    //     left: '3%',
    //     right: '4%',
    //     bottom: '3%',
    //     containLabel: true
    // },
        graphic:[{ // 一个图形元素，类型是 text，指定了 id。
            type: 'text',
            id: 'text2',
            style: {
            text: totalscore[totalscore.length - 1].toFixed(2),
               x: myChart.getZr().getWidth()*4/5,
            y: myChart.getZr().getHeight()/3, 
   textFill:'red',
//    textAlign: 'center', 
   textFont : 'bold 50px verdana'
    }
        },
        { // 一个图形元素，类型是 text，指定了 id。
            type: 'text',
            id: 'text0',
            style: {
            text: left.toFixed(1),
            x: myChart.getZr().getWidth()*1/5,
            y: myChart.getZr().getHeight()/3, 
   textFill:'green',
//    textAlign: 'center', 
   textFont : 'bold 50px verdana'
    }
        },
        ],


    xAxis : [
        {
            type : 'category',
            boundaryGap : false,
            data : datecategory
        }
    ],
    yAxis : [
        {
            type : 'value'
        }
    ],
    series : [
/*

    {
            name:'Make use of things you have learned',
            type:'line',
            stack: 'make',
            areaStyle: {},
            data:makeuseofthingsuhavelearned
        },

        // {
        //     name:'The total score of  number of tasks finished',
        //     type:'line',
        //     stack: 'number',
        //     areaStyle: {},
        //     data:thenumberoftasks_score
        // },




     
        {
            name:'Service for goal',
            type:'line',
            stack: 'serviceforgoal',
            areaStyle: {},
            data:serviceforgoal_score
        },




               {
            name:'Atom a difficult task',
            type:'line',
            stack: 'atomadifficulttask',
            areaStyle: {},
            data:atomadifficulttask_score
        },





            {
            name:'Only start a task',
            type:'line',
            stack: 'onlystartatask',
            areaStyle: {},
            data:onlystartatask_score
        },




    {
            name:'Finish difficult things',
            type:'line',
            stack: 'difficult',
            areaStyle: {},
            data:difficultthings
        },

    {
            name:'Threeminutes',
            type:'line',
            stack: 'three',
            areaStyle: {},
            data:threeminutes
        },

    {
            name:'Alwaysproft',
            type:'line',
            stack: 'three',
            areaStyle: {},
            data:alwaysprofit
        },


    {
            name:'Getlesson',
            type:'line',
            stack: 'getlesson',
            areaStyle: {},
            data:getlesson
        },

    




    {
            name:'Learntechuse',
            type:'line',
            stack: 'learntech',
            areaStyle: {},
            data:learntechuse
        },

*/
 
 /*    
  



    {
            name:'Start an important thing early',
            type:'line',
            stack: 'learntech',
            areaStyle: {},
            data:doanimportantthingearly
        },

    {
            name:'Mark a task immediately',
            type:'line',
            stack: 'learntech',
            areaStyle: {},
            data:markataskimmediately
        },




   {
            name:'Use brain to deal with things',
            type:'line',
            stack: 'brain',
            areaStyle: {},
            data:usebrain
        },

 {
            name:'Patience with the task',
            type:'line',
            stack: 'patience',
            areaStyle: {},
            data:patience
        },



  {
            name:'Battle with lower brain',
            type:'line',
            stack: 'battle',
            areaStyle: {},
            data:battlewithlowerbrain
        },



   {
            name:'Learn new things',
            type:'line',
            stack: 'learn',
            areaStyle: {},
            data:learnnewthings
        },

        {
            name:'Build frame and principle',
            type:'line',
            stack: 'learn',
            areaStyle: {},
            data:buildframeandprinciple
        },
        {
            name:'Accept fact ',
            type:'line',
            stack: 'learn',
            areaStyle: {},
            data:acceptfact
        },


        {
                name: 'MA5',
                type: 'line',
                color:'red',
                data: calculateMA(5, totalscore),
                smooth: true,
                lineStyle: {
                    normal: { opacity: 0.5 }
                }
            },

            {
                name: 'MA10',
                type: 'line',
                color:'black',
                data: calculateMA(10, totalscore),
                smooth: true,
                lineStyle: {
                    normal: { opacity: 0.5 }
                }
            },



*/

  //同时在图像中添加别的类型的曲线
       
        {
            name:'Executeability',
            type:'line',
            color:"brown",
  //          stack: 'e',
//            areaStyle: {},
            data:executeablity,
            smooth: true,
                lineStyle: {
                    normal: { opacity: 1.0, 
                    width:2,
                    
                    }

                }




        },


 {
            name:'todayscore',
            type:'line',
            // stack: 'e',
            color:'red',
            // areaStyle: {},
            data:todayscore,
            smooth: true,
                lineStyle: {
                    normal: { opacity: 1.0, 
                    width:10
                    }

                }


        },
        {
            name:'Totalscore',
            type:'line',
            smooth: true,
            color:'blue',
//            stack: 'e',
          
/*
  label: {
                normal: {
                    show: true,
                    position: 'top'
                }
            },
            areaStyle: {normal: {}},
*/  
                lineStyle: {
                    normal: { opacity: 1,
                    width:2
                    }

                },
          data:totalscore
        }
    ]
}
);
// 这里需要检查异步,必须要在图像再入完成后再
//myChart.getZr().dispose();
// var _zr = myChart.getZr();
// alert(_zr)
// alert(_zr.getWidth())
// alert(_zr.getHeight())
//  ShowObjProperty(_zr);
//     _zr.add(new echarts.graphic.Text({
//      style: {            
//    x: _zr.getWidth()*4/5,
//    y: _zr.getHeight()/3,
//    textFill:'red',
//    text:totalscore[totalscore.length - 1].toFixed(2),
// //    textAlign: 'center', 
//    textFont : 'bold 50px verdana'
//    }}  
//     ));
//     _zr.add(new echarts.graphic.Text({
//      style: {            
//    x: _zr.getWidth()*1/5,
//    y: _zr.getHeight()/3, 
//    textFill:'green',
//    text: '10000',
// //    textAlign: 'center', 
//    textFont : 'bold 50px verdana'
//    }}  
//     ));

}



loadlink(); // This will run on page load



setInterval(function(){
    loadlink() // this will run after every 5 seconds
}, 10*60000);


