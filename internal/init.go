// we can set a relative path to the local .hookup folder
// ex. git config --local core.hooksPath ./.hookup

// for worktree specify which worktree to look and update the hookPath to
// something like 'hu init --worktree [worktree repo]'
// they can leave blank in which the hu will look at all the available worktrees to find a .hookup folder

package internal

func Init() {
	// local repo folder name to hold git hooks
	// do i want to make this customizable
	const FOLDER = ".hookup"

	// check if there is a hookup folder and create if necessary

	// update the local git core.hookPath

}
