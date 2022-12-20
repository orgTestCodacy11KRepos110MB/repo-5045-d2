package e2etests

import (
	"testing"
)

func testRegression(t *testing.T) {
	tcs := []testCase{
		{
			name: "dagre_special_ids",
			script: `
ninety\nnine
eighty\reight
seventy\r\nseven
a\\yode -> there
a\\"ode -> there
a\\node -> there
`,
		},
		{
			name: "empty_sequence",
			script: `
A: hello {
  shape: sequence_diagram
}

B: goodbye {
  shape: sequence_diagram
}

A->B`,
		}, {
			name: "sequence_diagram_span_cover",
			script: `shape: sequence_diagram
b.1 -> b.1
b.1 -> b.1`,
		}, {
			name: "sequence_diagram_no_message",
			script: `shape: sequence_diagram
a: A
b: B`,
		},
		{
			name: "sequence_diagram_name_crash",
			script: `foo: {
	shape: sequence_diagram
	a -> b
}
foobar: {
	shape: sequence_diagram
	c -> d
}
foo -> foobar`,
		},
		{
			name: "sql_table_overflow",
			script: `
table: sql_table_overflow {
	shape: sql_table
	short: loooooooooooooooooooong
	loooooooooooooooooooong: short
}
table_constrained: sql_table_constrained_overflow {
	shape: sql_table
	short: loooooooooooooooooooong {
		constraint: unique
	}
	loooooooooooooooooooong: short {
		constraint: foreign_key
	}
}
`,
		},
		{
			name: "elk_alignment",
			script: `
direction: down

build_workflow: lambda-build.yaml {

	push: Push to main branch {
		style.font-size: 25
	}

	GHA: GitHub Actions {
		style.font-size: 25
	}

	S3.style.font-size: 25
	Terraform.style.font-size: 25
	AWS.style.font-size: 25

	push -> GHA: Triggers {
		style.font-size: 20
	}

	GHA -> S3: Builds zip and pushes it {
		style.font-size: 20
	}

	S3 <-> Terraform: Pulls zip to deploy {
		style.font-size: 20
	}

	Terraform -> AWS: Changes live lambdas {
		style.font-size: 20
	}
}

deploy_workflow: lambda-deploy.yaml {

	manual: Manual Trigger {
		style.font-size: 25
	}

	GHA: GitHub Actions {
		style.font-size: 25
	}

	AWS.style.font-size: 25

	Manual -> GHA: Launches {
		style.font-size: 20
	}

	GHA -> AWS: Builds zip\npushes them to S3.\n\nDeploys lambdas\nusing Terraform {
		style.font-size: 20
	}
}

apollo_workflow: apollo-deploy.yaml {

	apollo: Apollo Repo {
		style.font-size: 25
	}

	GHA: GitHub Actions {
		style.font-size: 25
	}

	AWS.style.font-size: 25

	apollo -> GHA: Triggered manually/push to master test test test test test test test {
		style.font-size: 20
	}

	GHA -> AWS: test {
		style.font-size: 20
	}
}
`,
		},
	}

	runa(t, tcs)
}
