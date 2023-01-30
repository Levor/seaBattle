# Sea battle game

### Purpose of this project: 
> The project was created to help teachers.


###  Opportunities:
1. Creation of an interactive game for several subjects at once and for a large number of topics.
2. Competitive component, which makes the lesson more interesting.
3. Preparation for the lesson, compiling and saving subjects, topics and questions in a json file.

## How the app works and what you need to know:
1. The game starts from the *.exe file.
2. At the first start, you need to launch the editor (File - Editor) and add a subject, topic and questions to the topic of the lesson (file "data.json" will be created).
3. After the item, topic and questions to the topic have been added, you can start a new game.
4. On subsequent launches, there is no need to add items.
5. When transferring the game to another PC, in order for the data (objects, topics, questions to be saved) you need to copy the file "data.json" 

## How packaging this game for differents systems

1. `go install fyne.io/fyne/v2/cmd/fyne@latest`
2. 
**Windows:**
`fyne package -os windows -icon logo.png`

**Linux**
`fyne package -os linux -icon logo.png`

**MacOS**
`fyne package -os darwin -icon logo.png`
