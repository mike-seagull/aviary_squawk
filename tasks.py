from invoke import task, Collection
from invoke.tasks import call
from os.path import join

PIPRUN="pipenv run"
PIPINST="pipenv install"
SRCDIR="src"
SRCFILEROOT="service"
BINDIR="bin"
BINFILE="squawk"
SRCMAIN=join(SRCDIR, "{}.py".format(SRCFILEROOT))
BINFILEFULLPATH=join(BINDIR,BINFILE)

@task
def clean(ctx):
    """ ...up after yourself """
    delete_files = [
        "dist/*",
        BINFILE,
        "{}.spec".format(SRCFILEROOT),
        "build/"
    ]
    for file in delete_files:
        ctx.run("rm -rf {}".format(file))

@task(help={'dev': 'dev dependencies'})
def install(ctx, dev=False):
    """  ...dependencies """
    if dev:
        ctx.run("{} --dev".format(PIPINST))
    else:
        ctx.run(PIPINST)

# run locally
def get_lambda_local_cmd():
    return '{} lambda invoke --event-file=test/event.json'.format(PIPRUN)
@task(name="run")
def run_lambda(ctx):
    """ runs aws lambda locally """
    ctx.run(get_lambda_local_cmd())
@task(help={'aws-lambda':"run aws lambda locally"})
def run(ctx, aws_lambda=False):
    """ ...locally """
    if aws_lambda:
        ctx.run(get_lambda_local_cmd())
    else:
        ctx.run('{} python {}'.format(PIPRUN, SRCMAIN))

# build
def get_lambda_package_cmd():
    return '{} lambda build'.format(PIPRUN)
@task(name="build")
def build_lambda(ctx):
    """ package for lambda """
    ctx.run(get_lambda_package_cmd())
@task(help={'aws-lambda':"build for aws lambda"})
def build(ctx, aws_lambda=False):
    """ package/bundle """
    if aws_lambda:
        ctx.run(get_lambda_package_cmd())
    else:
        ctx.run("{} pyinstaller --distpath {} --onefile {}".format(PIPRUN, BINDIR, SRCMAIN))
        ctx.run("mv {} {}".format(join(BINDIR, SRCFILEROOT), BINFILE))

ns = Collection()
ns.add_task(run)
ns.add_task(build)
ns.add_task(install)
ns.add_task(clean)
aws_lambda = Collection("lambda")
aws_lambda.add_task(run_lambda)
aws_lambda.add_task(build_lambda)
ns.add_collection(aws_lambda)
