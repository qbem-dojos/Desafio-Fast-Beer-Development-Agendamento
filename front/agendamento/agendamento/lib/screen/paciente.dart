import 'package:flutter/material.dart';
import 'package:flutter/src/widgets/container.dart';
import 'package:flutter/src/widgets/framework.dart';
import 'package:syncfusion_flutter_calendar/calendar.dart';

class PacienteScreen extends StatefulWidget {
  const PacienteScreen({super.key});

  @override
  State<PacienteScreen> createState() => _PacienteScreenState();
}

class _PacienteScreenState extends State<PacienteScreen> {
  @override
  Widget build(BuildContext context) {
    return SfCalendar(
          view: CalendarView.workWeek,
          dataSource: _getCalendarDataSource(),
          onTap: (calendarTapDetails) => DialogExample(),
          
          );
  }
}



class DataSource extends CalendarDataSource {
 DataSource(List<Appointment> source) {
   appointments = source;
 }
}


DataSource _getCalendarDataSource() {
   List<Appointment> appointments = <Appointment>[];
   appointments.add(
       Appointment(
         startTime: DateTime.now(),
         endTime: DateTime.now().add(
             const Duration(hours: 2)),
         isAllDay: false,
         subject: 'Meeting',
         color: Colors.blue,
         startTimeZone: 'America/Sao_Paulo',
         endTimeZone: 'America/Sao_Paulo'
       ));
  return DataSource(appointments);
}

class DialogExample extends StatelessWidget {

  @override
  Widget build(BuildContext context) {
    return TextButton(
      onPressed: () => showDialog<String>(
        context: context,
        builder: (BuildContext context) => AlertDialog(
          title: const Text('AlertDialog Title'),
          content: const Text('AlertDialog description'),
          actions: <Widget>[
            TextFormField(
            decoration: const InputDecoration(
            hintText: 'Enter your email',
            ),
            validator: (String? value) {
            if (value == null || value.isEmpty) {
            return 'Please enter some text';
            }
            return null;
            },
            ),
            TextButton(
              onPressed: () => Navigator.pop(context, 'Cancel'),
              child: const Text('Cancel'),
            ),
            TextButton(
              onPressed: () => Navigator.pop(context, 'OK'),
              child: const Text('OK'),
            ),
          ],
        ),
      ),
      child: const Text('Show Dialog'),
    );
  }
}
