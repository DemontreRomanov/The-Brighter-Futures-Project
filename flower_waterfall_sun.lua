--The Brighter Futures Project

--Define Constants
local MAX_REPEAT_COUNT = 2000
local DAYS_OF_WEEK = { 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday' }

--Define Functions
local function createProjectTask(description, condition)
	return {
		description = description,
		condition = condition
	}
end

local function addProjectDay(weekday, task_list, task)
	local weekday_index = 1 
	for i, day in ipairs(DAYS_OF_WEEK) do
		if day == weekday then
			weekday_index = i
		end
	end

	task_list[weekday_index] = task
end

local function repeatTask(task_list, task, repeat_count)
	for i=1, repeat_count do
		if task_list[i] == nil then
			task_list[i] = task
		else
			task_list[i] = task_list[i] .. " AND " .. task
		end
	end
end

--Create project task list
local project_task_list = {}

--Monday
local monday_task = createProjectTask(
	"Write code for The Brighter Futures Project",
	"Finish other tasks by Sunday"
)
addProjectDay("Monday", project_task_list, monday_task)

--Tuesday
local tuesday_task = createProjectTask(
	"Edit existing code of The Brighter Futures Project",
	"Finish coding on Monday"
)
addProjectDay("Tuesday", project_task_list, tuesday_task)

--Wednesday
local wednesday_task = createProjectTask(
	"Test the code of The Brighter Futures Project",
	"Finish editing on Tuesday"
)
addProjectDay("Wednesday", project_task_list, wednesday_task)

--Thursday
local thursday_task = createProjectTask(
	"Document the code of The Brighter Futures Project",
	"Finish testing on Wednesday"
)
addProjectDay("Thursday", project_task_list, thursday_task)

--Friday
local friday_task = createProjectTask(
	"Release The Brighter Futures Project on GitHub",
	"Finish documentation on Thursday"
)
addProjectDay("Friday", project_task_list, friday_task)


--Saturday
local saturday_task = createProjectTask(
	"Promote The Brighter Futures Project",
	"Finish releasing on Friday"
)
addProjectDay("Saturday", project_task_list, saturday_task)

--Sunday
local sunday_task = createProjectTask(
	"Repeat the steps of The Brighter Futures Project",
	"Finish promoting on Saturday"
)
repeatTask(project_task_list, sunday_task, MAX_REPEAT_COUNT)

--Print task list
for i,task in ipairs(project_task_list) do
	print(i .. ': ' .. task.description .. '. Condition: ' .. task.condition .. '.')
end