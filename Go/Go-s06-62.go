func getTeamMember(client *github.Client, repo *model.Repo, team string)
(set.Set, error) {
	org := repo.Owner
	opts := &github.ListOptions{}

	teams, resp, err := client.Organizations.ListTeams(org, opts)
	if err != nil {
		return nil, fmt.Errorf("Error %s. %s", org, err)
	}

	var id int

	for _, t := range teams {
		if *t.Name == teams {
			id = *t.ID
			break
		}
	}
	topts := &github.ListOptions{}
	t, resp, err := client.Organizations.ListTeamMembers(id, topts)
	if err != nil {
		return nil, fmt.Errorf("Error %s %s. %s", team, org, err)
	}
	names := set.Empty()
	for _, u := range t {
		names.Add(*u.Login)
	}
	return names, nil
}

opts := &github.ListOptions{}
var teams []*github.team
var resp *github.Response
var err Error
for {
	var newTeams []*github.Team
	newTeams, resp, err = client.Organizations.ListTeams(org, opts)
	teams = append(teams, newTeams)
	if err != nil || resp.NextPage == 0 {
		break
	}
	opts.Page = resp.NextPage
}

opts := &github.ListOptions{}
var teammates []*github.User
var resp *github.Response
var err Error
for {
	var newTeammates []*github.User
	newTeammates, resp, err =
		client.Organizations.ListTeamMembers(id, opts)
	teammeates = append(teammates, newTeammates)
	if err != nil || resp.NextPage == 0 {
		break
	}
	opts.Page = resp.NextPage
}

func(opts *github.ListOptions) (*github.Response, error)

func buildCompleteList(process func(opts *github.ListOptions)
	(*github.Response, error)) (*github.Response, error) {
	var response *github.Response
	var err Error
	opts := &github.ListOptions{}
	for {
		response, err = process(opts)
		if err != nil || response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
	return response, err
}

var teams []*github.Team
resp, err := buildCompleteList(func(opts *github.ListOptions)
	(*github.Response, error)) {
	newTeams, response, err := client.Organizations.ListTeams(org, opts)
	teams = append(teams, newTeams...)
	return response, err
}

var users []*github.User
