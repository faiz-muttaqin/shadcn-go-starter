import { createFileRoute, Link, Outlet } from '@tanstack/react-router'
import { LearnMore } from '@/components/LearnMore'
import { cn } from '@/lib/utils'
import dashboardDark from '@/features/auth/sign-in/assets/dashboard-dark.png'
import dashboardLight from '@/features/auth/sign-in/assets/dashboard-light.png'

export const Route = createFileRoute('/clerk/(auth)')({
  component: ClerkAuthLayout,
})

function ClerkAuthLayout() {
  return (
    <div className='relative container grid h-svh flex-col items-center justify-center lg:max-w-none lg:grid-cols-2 lg:px-0'>
      <div className='lg:p-8'>
        <div className='relative mx-auto flex w-full flex-col items-center justify-center gap-4'>
          <LearnMore
            defaultOpen
            triggerProps={{
              className: 'absolute -top-12 end-0 sm:end-20 size-6',
            }}
            contentProps={{ side: 'top', align: 'end', className: 'w-auto' }}
          >
            Welcome to the example Clerk auth page. <br />
            Back to{' '}
            <Link
              to='/'
              className='underline decoration-dashed underline-offset-2'
            >
              Dashboard
            </Link>{' '}
            ?
          </LearnMore>
          <Outlet />
        </div>
      </div>
      {/* <div className='bg-muted relative hidden h-full flex-col p-10 text-white lg:flex dark:border-e'>
        <div className='absolute inset-0 bg-slate-500' />
        <Link
          to='/'
          className='relative z-20 flex items-center text-lg font-medium'
        >
          <Logo className='me-2' />
          Shadcn Admin Go Starter
        </Link>

        <ClerkFullLogo className='relative m-auto size-96' />

        <div className='relative z-20 mt-auto'>
          <blockquote className='space-y-2'>
            <p className='text-lg'>
              &ldquo; Lorem ipsum dolor sit amet consectetur adipisicing elit.
              Sint, magni debitis inventore asperiores velit! &rdquo;
            </p>
            <footer className='text-sm'>John Doe</footer>
          </blockquote>
        </div>
      </div> */}
      
      <div
        className={cn(
          'bg-muted relative h-full overflow-hidden max-lg:hidden',
          '[&>img]:absolute [&>img]:top-[15%] [&>img]:left-20 [&>img]:h-full [&>img]:w-full [&>img]:object-cover [&>img]:object-top-left [&>img]:select-none'
        )}
      >
        <img
          src={dashboardLight}
          className='dark:hidden'
          width={1024}
          height={1151}
          alt='Shadcn-Admin'
        />
        <img
          src={dashboardDark}
          className='hidden dark:block'
          width={1024}
          height={1138}
          alt='Shadcn-Admin'
        />
      </div>
    </div>
  )
}
