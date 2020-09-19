package main

import (
	"flag"
	"fmt"
	"github.com/libgit2/git2go/v30"
)
import "log"

func main() {
	repoPath := flag.String("repo", "/Users/daniel/learn/rust-by-leetcode", "path to the git repository")
	flag.Parse()

	repo, err := git.OpenRepository(*repoPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(repo)
	defer repo.Free()
	//odb, err := repo.Odb()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer odb.Free()
	//odb.ForEach(func(oid *git.Oid) error {
	//	obj, err := repo.Lookup(oid)
	//	if err != nil {
	//		log.Fatal("Lookup:", err)
	//	}
	//	//x := obj.Type()
	//	//log.Println(x)
	//	switch obj.Type() {
	//	default:
	//	case git.ObjectBlob:
	//		blob, _ := obj.AsBlob()
	//		fmt.Printf("==================================\n")
	//		fmt.Printf("Type: %s\n", blob.Type())
	//		fmt.Printf("Id:   %s\n", blob.Id())
	//		fmt.Printf("Size: %d\n", blob.Size())
	//	case git.ObjectCommit:
	//		commit, _ := obj.AsCommit()
	//		fmt.Printf("==================================\n")
	//		fmt.Printf("Type: %s\n", commit.Type())
	//		fmt.Printf("Id:   %s\n", commit.Id())
	//		author := commit.Author()
	//		fmt.Printf("    Author:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", author.Name, author.Email, author.When)
	//		committer := commit.Committer()
	//		fmt.Printf("    Committer:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", committer.Name, committer.Email, committer.When)
	//		fmt.Printf("    ParentCount: %d\n", commit.ParentCount())
	//		fmt.Printf("    TreeId:      %s\n", commit.TreeId())
	//		fmt.Printf("    Message:\n\n        %s\n\n", strings.Replace(commit.Message(), "\n", "\n        ", -1))
	//	case git.ObjectTree:
	//		tree, _ := obj.AsTree()
	//		fmt.Printf("==================================\n")
	//		fmt.Printf("Type: %s\n", tree.Type())
	//		fmt.Printf("Id:   %s\n", tree.Id())
	//		fmt.Printf("    EntryCount: %d\n", tree.EntryCount())
	//	case git.ObjectTag:
	//		tag, _ := obj.AsTag()
	//		fmt.Printf("==================================\n")
	//		fmt.Printf("Type: %s\n", tag.Type())
	//		fmt.Printf("Id:   %s\n", tag.Id())
	//		fmt.Printf("Id:   %s\n", tag.Name())
	//	}
	//	return nil
	//})
	opt := git.DiffOptions{}
	walk, err := repo.Walk()
	if err != nil {
		log.Fatal(err)
	}
	defer walk.Free()
	_ = walk.PushHead()
	_ = walk.Iterate(func(commit *git.Commit) bool {
		//fmt.Printf("%s\n", commit)
		//sid, _ :=commit.ShortId()
		//		fmt.Printf("%s\n", sid)
		commitTime := commit.Author().When
		fmt.Printf("%s\n", commitTime.Format("Mon Jan 2 15:04:05 -0700 MST 2006"))
		id := commit.Id()
		if commit.ParentCount() >= 1 {
			pid := commit.Parent(0).Id()
			o, _ := repo.RevparseSingle(id.String())
			t, _ := o.Peel(git.ObjectTree)
			t1, _ := t.AsTree()
			//			fmt.Printf("%s\n", t)
			po, _ := repo.RevparseSingle(pid.String())
			pt, _ := po.Peel(git.ObjectTree)
			pt1, _ := pt.AsTree()
			diff, _ := repo.DiffTreeToTree(pt1, t1, &opt)
			defer diff.Free()
			//			fmt.Printf("%s\n", diff)
			stats, _ := diff.Stats()
			//			fmt.Printf("%s\n", stats)
			stat, _ := stats.String(git.DiffStatsFull, 40)
			fmt.Printf("%s\n", stat)
		}
		return true
	})
}
