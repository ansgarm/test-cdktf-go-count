package main

import (
	"cdk.tf/go/stack/generated/random"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	random.NewRandomProvider(stack, jsii.String("random"), &random.RandomProviderConfig{})

	pets := random.NewPet(stack, jsii.String("pet"), &random.PetConfig{
		Count: 2,
	})

	cdktf.NewTerraformOutput(stack, jsii.String("first"), &cdktf.TerraformOutputConfig{
		Value: jsii.String("${" + *pets.Fqn() + "[0].id}"),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "test-cdktf-go-count")

	app.Synth()
}
